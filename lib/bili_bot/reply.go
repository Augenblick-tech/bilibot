package bilibot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Augenblick-tech/bilibot/pkg/e"
	"github.com/Augenblick-tech/bilibot/pkg/utils"
	"github.com/spf13/viper"
)

type ReplyResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		SuccessToast string      `json:"success_toast"`
		Emote        interface{} `json:"emote"`
	} `json:"data"`
}

func AddReply(cookies string, typeID int, oid string, message string) (*ReplyResponse, error) {
	cookie := utils.StrToMap(cookies)
	url := "http://api.bilibili.com/x/v2/reply/add"
	client := &http.Client{}
	req, err := http.NewRequest(
		"POST",
		url,
		strings.NewReader(
			fmt.Sprintf("plat=1&type=%d&oid=%s&message=%s&csrf=%s", typeID, oid, message, cookie["bili_jct"]),
		),
	)
	if err != nil {
		return nil, err
	}
	req.AddCookie(&http.Cookie{Name: "SESSDATA", Value: cookie["SESSDATA"]})
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", viper.GetString("server.user_agent"))
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var ReplyResponse ReplyResponse
	err = json.Unmarshal(body, &ReplyResponse)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return &ReplyResponse, e.ERR_COMMENT_REPLY_FAIL
	}
	return &ReplyResponse, nil
}
