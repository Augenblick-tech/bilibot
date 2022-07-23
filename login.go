package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/lonzzi/BiliUpDynamicBot/e"
)

type QRCodeResponse struct {
	Code int `json:"code"`
	TS   int `json:"ts"`
	Data struct {
		Url      string `json:"url"`      // 二维码内容url
		OauthKey string `json:"oauthKey"` // 扫码登录秘钥
	} `json:"data"`
}

type LoginResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	TS      int         `json:"ts"`
	Status  bool        `json:"status"`
	Data    interface{} `json:"data"`
}

func Login() (*AccountInfo, error) {
	qrCodeBody, err := Fetch("http://passport.bilibili.com/qrcode/getLoginUrl")
	if err != nil {
		return nil, err
	}
	var qrCodeResponse QRCodeResponse
	err = json.Unmarshal(qrCodeBody, &qrCodeResponse)
	if err != nil {
		return nil, err
	}

	fmt.Println("QR Code: ", qrCodeResponse.Data.Url)

	client := &http.Client{}
	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()
	cnt := 1
	for {
		<-ticker.C
		cnt++
		if cnt > 10 {
			break
		}
		loginReq, err := http.NewRequest(
			"POST",
			"http://passport.bilibili.com/qrcode/getLoginInfo",
			strings.NewReader(fmt.Sprintf("oauthKey=%s", qrCodeResponse.Data.OauthKey)),
		)
		if err != nil {
			log.Fatal(err)
		}
		loginReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		loginReq.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36")

		loginResp, err := client.Do(loginReq)
		if err != nil {
			log.Fatal(err)
		}

		loginBody, err := ioutil.ReadAll(loginResp.Body)
		if err != nil {
			log.Fatal(err)
		}
		var loginResponse LoginResponse
		err = json.Unmarshal(loginBody, &loginResponse)
		if err != nil {
			log.Fatal(err)
		}

		loginResp.Body.Close()

		if v, ok := loginResponse.Data.(float64); ok {
			switch v {
			case -1:
				log.Println(e.KEY_INVALID)
			case -2:
				log.Println(e.KEY_TIMEOUT)
			case -4:
				log.Println(e.NOT_SCAN)
			case -5:
				log.Println(e.NOT_CONFIRM)
			default:
				log.Println("未知代码", v)
			}
		}

		if loginResponse.Status {
			log.Println("登录成功")
			url := loginResponse.Data.(map[string]interface{})["url"].(string)
			params := strings.Split(url, "&")
			accountData := StrUrl2Map(params)
			return &AccountInfo{
				SESSDATA: accountData["SESSDATA"],
				BiliJct:  accountData["bili_jct"],
			}, nil
		}
	}
	return nil, e.ERR_LOGIN_FAIL
}
