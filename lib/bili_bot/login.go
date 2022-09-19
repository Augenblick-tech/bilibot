package bilibot

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Augenblick-tech/bilibot/pkg/client"
	"github.com/Augenblick-tech/bilibot/pkg/e"
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
	Mid  uint   `json:"mid"`
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
	var qrCodeResponse QRCodeResponse
	URL := "http://passport.bilibili.com/qrcode/getLoginUrl"

	v := client.NewVisitor()
	v.OnResponse(func(r *client.Response) {
		json.Unmarshal(r.Body, &qrCodeResponse)
	})

	return &qrCodeResponse, v.Visit(URL)
}

func GetLoginInfo(oauthKey string) ([]*http.Cookie, error) {
	var loginResponse LoginResponse
	URL := "http://passport.bilibili.com/qrcode/getLoginInfo"

	v := client.NewVisitor()
	v.OnResponse(func(r *client.Response) {
		json.Unmarshal(r.Body, &loginResponse)
	})

	v.Post(URL, []byte(fmt.Sprintf("oauthKey=%s", oauthKey)))

	if v, ok := loginResponse.Data.(float64); ok {
		switch v {
		case -1:
			return nil, e.KeyInvalid
		case -2:
			return nil, e.KeyTimeout
		case -4:
			return nil, e.Waiting
		case -5:
			return nil, e.NotConfirmed
		default:
			return nil, e.ErrBiliUndefined
		}
	}

	if loginResponse.Status {
		return v.Cookies(URL), nil
	}

	return nil, e.ErrLoginFailed
}

func GetInfo(mid string) (*AuthorInfo, error) {
	var authorInfo AuthorInfo
	URL := fmt.Sprintf("http://api.bilibili.com/x/space/acc/info?mid=%s", mid)

	v := client.NewVisitor()
	v.OnResponse(func(r *client.Response) {
		json.Unmarshal(r.Body, &authorInfo)
	})

	if err := v.Visit(URL); err != nil {
		return nil, err
	}

	if authorInfo.Code == -404 {
		return nil, e.ErrNotFound
	}

	return &authorInfo, nil
}

func GetBotInfo(cookie *http.Cookie) (*BotInfo, error) {
	var botInfo BotInfo
	URL := "http://api.bilibili.com/nav"

	v := client.NewVisitor()
	v.SetCookies(URL, []*http.Cookie{
		cookie,
	})
	v.OnResponse(func(r *client.Response) {
		json.Unmarshal(r.Body, &botInfo)
	})

	return &botInfo, v.Visit("http://api.bilibili.com/nav")
}
