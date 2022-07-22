package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/lonzzi/BiliUpDynamicBot/e"
	"github.com/spf13/viper"
)

type Dynamic struct {
	Code int `json:"code"`
	Data struct {
		HasMore bool `json:"has_more"`
		Items   []struct {
			IDStr   string `json:"id_str"` // 动态ID
			Modules struct {
				Author  Author  `json:"module_author"`  // 动态作者
				Content Content `json:"module_dynamic"` // 动态内容
			} `json:"modules"`
		} `json:"items"`
	} `json:"data"`
}

type Author struct {
	Mid       int    `json:"mid"`    // 作者ID
	Name      string `json:"name"`   // 作者名称
	Face      string `json:"face"`   // 作者头像
	TimeStamp int64  `json:"pub_ts"` // 作者发布时间
}

type Content struct {
	Desc struct {
		Text string `json:"text"` // 动态内容
	} `json:"desc"`
}

type BriefDynamic struct {
	IDStr     string  // 动态ID
	Author    Author  // 动态作者
	Content   Content // 动态内容
}

func fetch(url string) string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Fetch Error: ", err)
		return ""
	}
	if resp.StatusCode != http.StatusOK {
		log.Println("Error: ", resp.StatusCode)
		return ""
	}

	defer resp.Body.Close()

	log.Println(resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(body)
}

func GetLatestDynamic(mid string) *BriefDynamic {
	url := "https://api.bilibili.com/x/polymer/web-dynamic/v1/feed/space?host_mid=" + mid
	body := fetch(url)
	var dynamic Dynamic
	json.Unmarshal([]byte(body), &dynamic)
	return &BriefDynamic{
		IDStr:   dynamic.Data.Items[0].IDStr,
		Author:  dynamic.Data.Items[0].Modules.Author,
		Content: dynamic.Data.Items[0].Modules.Content,
	}
}

func CommentReply(typeID int, oid string, message string) (string, error) {
	cookie := "SESSDATA=" + viper.GetString("account.SESSDATA")
	url := "http://api.bilibili.com/x/v2/reply/add"
	client := &http.Client{}
	req, err := http.NewRequest(
		"POST",
		url,
		strings.NewReader(
			fmt.Sprintf("plat=1&type=%d&message=%s&oid=%s&csrf=%s", typeID, message, oid, viper.GetString("account.bili_jct")),
		),
	)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Cookie", cookie)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return string(body), errors.New("评论错误")
	}
	return string(body), nil
}

func DynamicReply(dynamic BriefDynamic, message string) (string, error) {
	return CommentReply(e.DynamicCommentCode, dynamic.IDStr, message)
}