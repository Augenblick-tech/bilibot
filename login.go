package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
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

func Login() *AccountInfo {
	client := &http.Client{}

	getQRCodeReq, err := http.NewRequest("GET", "http://passport.bilibili.com/qrcode/getLoginUrl", nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	getQRCodeReq.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36")
	getQRCodeResp, err := client.Do(getQRCodeReq)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer getQRCodeResp.Body.Close()

	var qrCodeResponse QRCodeResponse
	qrCodeBody, err := ioutil.ReadAll(getQRCodeResp.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(qrCodeBody, &qrCodeResponse)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	fmt.Println("QR Code: ", qrCodeResponse.Data.Url)

	ticker := time.NewTicker(time.Second * 5)
	cnt := 1
	defer ticker.Stop()
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
			panic(err)
		}
		loginReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		loginReq.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36")

		loginResp, err := client.Do(loginReq)
		if err != nil {
			panic(err)
		}
		defer loginResp.Body.Close()

		loginBody, err := ioutil.ReadAll(loginResp.Body)
		if err != nil {
			panic(err)
		}
		var loginResponse LoginResponse
		err = json.Unmarshal(loginBody, &loginResponse)
		if err != nil {
			panic(err)
		}
		fmt.Println("Login Response: ", loginResponse)
		if loginResponse.Status {
			url := loginResponse.Data.(map[string]interface{})["url"].(string)
			params := strings.Split(url, "&")
			accountData := StrUrl2Map(params)
			return &AccountInfo{
				SESSDATA: accountData["SESSDATA"],
				BiliJct:  accountData["bili_jct"],
			}
		}
	}
	return nil
}

