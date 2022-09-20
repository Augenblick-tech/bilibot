package bilibot

import (
	"encoding/json"
	"fmt"

	"github.com/Augenblick-tech/bilibot/pkg/client"
	"github.com/Augenblick-tech/bilibot/pkg/utils"
)

type ReplyResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		SuccessToast string      `json:"success_toast"`
		Emote        interface{} `json:"emote"`
	} `json:"data"`
}

func AddReply(cookie string, typeID int, oid string, message string) (*ReplyResponse, error) {
	var ReplyResponse ReplyResponse
	URL := "http://api.bilibili.com/x/v2/reply/add"
	cookies := utils.StrToCookies(cookie)

	v := client.NewVisitor()
	v.SetCookies(URL, cookies)
	v.OnResponse(func(r *client.Response) {
		json.Unmarshal(r.Body, &ReplyResponse)
	})

	cookieMap := utils.CookieToMap(cookies)

	v.Post(URL, []byte(fmt.Sprintf("plat=1&type=%d&oid=%s&message=%s&csrf=%s", typeID, oid, message, cookieMap["bili_jct"])))

	return &ReplyResponse, nil
}
