package model

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/lonzzi/BiliUpDynamicBot/pkg/e"
	"github.com/lonzzi/BiliUpDynamicBot/pkg/utils"
	"github.com/spf13/viper"
)

type DynamicObj struct {
	Code int `json:"code"`
	Data struct {
		HasMore bool      `json:"has_more"`
		Items   []Dynamic `json:"items"`
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

type Dynamic struct {
	IDStr   string `json:"id_str"` // 动态ID
	Modules struct {
		Author  Author  `json:"module_author"`  // 动态作者
		Content Content `json:"module_dynamic"` // 动态内容
	} `json:"modules"`
}

func GetDynamic(mid string) ([]Dynamic, error) {
	url := "https://api.bilibili.com/x/polymer/web-dynamic/v1/feed/space?host_mid=" + mid
	body, err := utils.Fetch(url)
	if err != nil {
		return nil, err
	}
	var dynamicObj DynamicObj
	json.Unmarshal(body, &dynamicObj)
	return dynamicObj.Data.Items, nil
}

func IsDynamicExist(dynamics []Dynamic, dynamic Dynamic) bool {
	for _, v := range dynamics {
		if v.IDStr == dynamic.IDStr {
			return true
		}
	}
	return false
}

func DynamicReply(dynamic Dynamic, message string) (*ReplyResponse, error) {
	ReplyResponse, err := AddReply(string(rune(e.DynamicCommentCode)), dynamic.IDStr, message)
	if err != nil {
		return nil, e.ERR_REPLY_DYNAMIC
	}

	return ReplyResponse, nil
}

func GetHistoryDynamics(filePath string) ([]Dynamic, error) {
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
	var oldDynamics []Dynamic
	if err := json.Unmarshal(fileContent, &oldDynamics); err != nil && len(fileContent) != 0 {
		return nil, e.ERR_UNMARSHAL
	}
	return oldDynamics, nil
}

func AddNewDynamic(oldDynamics []Dynamic, dynamic Dynamic) ([]Dynamic, error) {
	if IsDynamicExist(oldDynamics, dynamic) {
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
