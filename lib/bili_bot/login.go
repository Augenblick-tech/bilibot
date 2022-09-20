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
		Url       string `json:"url"`        // 二维码内容url
		QrcodeKey string `json:"qrcode_key"` // 扫码登录秘钥
	} `json:"data"`
}

type LoginResponse struct {
	Code    int       `json:"code"`
	Message string    `json:"message"`
	Data    LoginData `json:"data"`
}

type LoginData struct {
	RefreshToken string `json:"refresh_token"`
	TimeStamp    int    `json:"timestamp"`
	Code         int    `json:"code"`
	Message      string `json:"message"`
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
	URL := "https://passport.bilibili.com/x/passport-login/web/qrcode/generate"

	v := client.NewVisitor()
	v.OnResponse(func(r *client.Response) {
		json.Unmarshal(r.Body, &qrCodeResponse)
	})

	return &qrCodeResponse, v.Visit(URL)
}

func GetLoginInfo(QrcodeKey string) ([]*http.Cookie, error) {
	var loginResponse LoginResponse
	URL := fmt.Sprintf("https://passport.bilibili.com/x/passport-login/web/qrcode/poll?qrcode_key=%s", QrcodeKey)

	v := client.NewVisitor()
	v.OnResponse(func(r *client.Response) {
		json.Unmarshal(r.Body, &loginResponse)
	})

	v.Visit(URL)

	return v.Cookies(URL), nil
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
