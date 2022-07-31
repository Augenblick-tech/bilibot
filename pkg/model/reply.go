package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/lonzzi/bilibot/pkg/e"
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

func AddReply(typeID string, oid string, message string) (*ReplyResponse, error) {
	cookie := "SESSDATA=" + viper.GetString("account.SESSDATA")
	url := "http://api.bilibili.com/x/v2/reply/add"
	client := &http.Client{}
	req, err := http.NewRequest(
		"POST",
		url,
		strings.NewReader(
			fmt.Sprintf("plat=1&type=%s&oid=%s&message=%s&csrf=%s", typeID, oid, message, viper.GetString("account.bili_jct")),
		),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Cookie", cookie)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
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
