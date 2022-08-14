package bilibot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"strings"

	"github.com/Augenblick-tech/bilibot/pkg/e"
	"github.com/Augenblick-tech/bilibot/pkg/utils"
	"github.com/spf13/viper"
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

type AuthorInfo struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    AuthorData `json:"data"`
}

type AuthorData struct {
	Mid  int    `json:"mid"`
	Name string `json:"name"`
	Sex  string `json:"sex"`
	Face string `json:"face"`
}

type BotInfo struct {
	Code    int     `json:"code"`
	Message string  `json:"message"`
	Data    BotData `json:"data"`
}

type BotData struct {
	Mid     uint   `json:"mid"`
	Name    string `json:"uname"`
	IsLogin bool   `json:"isLogin"`
	Face    string `json:"face"`
}

func GetLoginUrl() (*QRCodeResponse, error) {
	qrCodeBody, err := utils.Fetch("http://passport.bilibili.com/qrcode/getLoginUrl")
	if err != nil {
		return nil, err
	}
	var qrCodeResponse QRCodeResponse
	return &qrCodeResponse, json.Unmarshal(qrCodeBody, &qrCodeResponse)
}

func GetLoginInfo(oauthKey string) ([]*http.Cookie, error) {
	client := &http.Client{}
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	client.Jar = jar

	loginReq, err := http.NewRequest(
		"POST",
		"http://passport.bilibili.com/qrcode/getLoginInfo",
		strings.NewReader(fmt.Sprintf("oauthKey=%s", oauthKey)),
	)
	if err != nil {
		return nil, err
	}
	loginReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	loginReq.Header.Set("User-Agent", viper.GetString("server.user_agent"))

	loginResp, err := client.Do(loginReq)
	if err != nil {
		return nil, err
	}

	loginBody, err := io.ReadAll(loginResp.Body)
	if err != nil {
		return nil, err
	}

	var loginResponse LoginResponse
	err = json.Unmarshal(loginBody, &loginResponse)
	if err != nil {
		return nil, err
	}

	defer loginResp.Body.Close()

	if v, ok := loginResponse.Data.(float64); ok {
		switch v {
		case -1:
			return nil, e.KEY_INVALID
		case -2:
			return nil, e.KEY_TIMEOUT
		case -4:
			return nil, e.NOT_SCAN
		case -5:
			return nil, e.NOT_CONFIRM
		}
	}

	if loginResponse.Status {
		return client.Jar.Cookies(loginReq.URL), nil
	}

	return nil, e.ERR_LOGIN_FAIL
}

func GetInfo(mid string) (*AuthorInfo, error) {
	body, err := utils.Fetch(fmt.Sprintf("http://api.bilibili.com/x/space/acc/info?mid=%s", mid))
	if err != nil {
		return nil, err
	}
	var botInfo AuthorInfo
	return &botInfo, json.Unmarshal(body, &botInfo)
}

func GetBotInfo(cookie *http.Cookie) (*BotInfo, error) {
	body, err := utils.Fetch("http://api.bilibili.com/nav", cookie)
	if err != nil {
		return nil, err
	}
	var botInfo BotInfo
	return &botInfo, json.Unmarshal(body, &botInfo)
}
