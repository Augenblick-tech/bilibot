package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/lonzzi/BiliUpDynamicBot/pkg/e"
	"github.com/lonzzi/BiliUpDynamicBot/pkg/utils"
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
	IDStr   string  // 动态ID
	Author  Author  // 动态作者
	Content Content // 动态内容
}

type CommentResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		SuccessToast string      `json:"success_toast"`
		Emote        interface{} `json:"emote"`
	} `json:"data"`
}

func GetLatestDynamic(mid string) (*BriefDynamic, error) {
	url := "https://api.bilibili.com/x/polymer/web-dynamic/v1/feed/space?host_mid=" + mid
	body, err := utils.Fetch(url)
	if err != nil {
		return nil, err
	}
	var dynamic Dynamic
	json.Unmarshal(body, &dynamic)
	return &BriefDynamic{
		IDStr:   dynamic.Data.Items[0].IDStr,
		Author:  dynamic.Data.Items[0].Modules.Author,
		Content: dynamic.Data.Items[0].Modules.Content,
	}, nil
}

func CommentReply(typeID int, oid string, message string) (*CommentResponse, error) {
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
		log.Fatal(err)
	}
	req.Header.Set("Cookie", cookie)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var commentResponse CommentResponse
	err = json.Unmarshal(body, &commentResponse)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return &commentResponse, e.ERR_COMMENT_REPLY_FAIL
	}
	return &commentResponse, nil
}

func DynamicReply(dynamic BriefDynamic, message string) (*CommentResponse, error) {
	return CommentReply(e.DynamicCommentCode, dynamic.IDStr, message)
}

func MakeReply(oldDynamics []BriefDynamic, dynamic BriefDynamic, message string) (*CommentResponse, error) {
	commentResponse, err := DynamicReply(dynamic, message)
	if err != nil {
		return nil, e.ERR_REPLY_DYNAMIC
	}

	return commentResponse, nil
}

func GetHistoryDynamics(filePath string) ([]BriefDynamic, error) {
	// 读出历史数据
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			// 如果文件不存在，则创建文件
			err = ioutil.WriteFile(filePath, []byte("[]"), 0666)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, e.ERR_READFILE
	}
	var oldDynamics []BriefDynamic
	if err := json.Unmarshal(fileContent, &oldDynamics); err != nil && len(fileContent) != 0 {
		return nil, e.ERR_UNMARSHAL
	}
	return oldDynamics, nil
}

func IsExistDynamic(dynamics []BriefDynamic, dynamic BriefDynamic) bool {
	for _, v := range dynamics {
		if v.IDStr == dynamic.IDStr {
			return true
		}
	}
	return false
}

func AddNewDynamic(oldDynamics []BriefDynamic, dynamic BriefDynamic) ([]BriefDynamic, error) {
	if IsExistDynamic(oldDynamics, dynamic) {
		return oldDynamics, e.ERR_DYNAMIC_EXIST
	}

	oldDynamics = append(oldDynamics, dynamic)

	// 历史动态只保留10个
	if len(oldDynamics) > 10 {
		oldDynamics = oldDynamics[len(oldDynamics)-10:]
	}

	// 写入新数据
	fileContent, err := json.Marshal(oldDynamics)
	if err != nil {
		return nil, e.ERR_MARSHAL
	}
	if err := ioutil.WriteFile(viper.GetString("HistroyDynamicPath"), fileContent, 0644); err != nil {
		return nil, e.ERR_WRITEFILE
	}

	return oldDynamics, nil
}
