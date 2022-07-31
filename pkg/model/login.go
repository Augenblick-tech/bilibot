package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"strings"
	"time"

	"github.com/lonzzi/bilibot/pkg/e"
	"github.com/lonzzi/bilibot/pkg/utils"
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

type LoginInfo struct {
	UID      string
	SESSDATA string
	BiliJct  string
}

func GetLoginUrl() (*QRCodeResponse, error) {
	qrCodeBody, err := utils.Fetch("http://passport.bilibili.com/qrcode/getLoginUrl")
	if err != nil {
		return nil, err
	}
	var qrCodeResponse QRCodeResponse
	err = json.Unmarshal(qrCodeBody, &qrCodeResponse)
	if err != nil {
		return nil, err
	}

	return &qrCodeResponse, nil
}

func GetLoginInfo(oauthKey string, timeout int) (*LoginInfo, error) {
	client := &http.Client{}
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	client.Jar = jar
	interval := 3
	var finalErr error
	ticker := time.NewTicker(time.Second * time.Duration(interval))
	defer ticker.Stop()
	cnt := 0
	for {
		<-ticker.C
		cnt++
		if cnt > timeout/interval {
			break
		}
		loginReq, err := http.NewRequest(
			"POST",
			"http://passport.bilibili.com/qrcode/getLoginInfo",
			strings.NewReader(fmt.Sprintf("oauthKey=%s", oauthKey)),
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

		defer loginResp.Body.Close()

		if v, ok := loginResponse.Data.(float64); ok {
			switch v {
			case -1:
				log.Println(e.KEY_INVALID)
				finalErr = e.KEY_INVALID
			case -2:
				log.Println(e.KEY_TIMEOUT)
				finalErr = e.KEY_TIMEOUT
			case -4:
				log.Println(e.NOT_SCAN)
				finalErr = e.NOT_SCAN
			case -5:
				log.Println(e.NOT_CONFIRM)
				finalErr = e.NOT_CONFIRM
			default:
				log.Println("未知代码", v)
			}
		}

		if loginResponse.Status {
			log.Println("登录成功")
			url := loginResponse.Data.(map[string]interface{})["url"].(string)
			params := strings.Split(url, "&")
			accountData := utils.StrUrlToMap(params)
			return &LoginInfo{
				UID:      accountData["DedeUserID"],
				SESSDATA: accountData["SESSDATA"],
				BiliJct:  accountData["bili_jct"],
			}, nil
		}
	}

	return nil, finalErr
}
