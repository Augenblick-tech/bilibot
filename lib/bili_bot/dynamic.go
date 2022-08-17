package bilibot

import (
	"encoding/json"

	"github.com/Augenblick-tech/bilibot/pkg/utils"
)

type Dynamics struct {
	Code int `json:"code"`
	Data struct {
		HasMore bool      `json:"has_more"` // 是否还有更多动态
		Items   []Dynamic `json:"items"`
		Offset  string    `json:"offset"` // 动态偏移量，触发下一页动态
	} `json:"data"`
}

type Author struct {
	Mid   uint    `json:"mid"`    // 作者ID
	Name  string `json:"name"`   // 作者名称
	Face  string `json:"face"`   // 作者头像
	PubTS uint64  `json:"pub_ts"` // 作者发布时间
}

type Content struct {
	Desc struct {
		Text string `json:"text"` // 动态内容
	} `json:"desc"`
}

type Dynamic struct {
	ID      string `json:"id_str"` // 动态ID
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

	var dynamics Dynamics
	err = json.Unmarshal(body, &dynamics)

	return dynamics.Data.Items, err
}
