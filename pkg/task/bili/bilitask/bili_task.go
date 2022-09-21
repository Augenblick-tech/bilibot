package bilitask

import (
	"fmt"
	"log"

	"github.com/Augenblick-tech/bilibot/lib/bili_bot"
	"github.com/Augenblick-tech/bilibot/pkg/email"
	"github.com/Augenblick-tech/bilibot/pkg/plugin"
	"github.com/Augenblick-tech/bilibot/pkg/services/bot"
	"github.com/Augenblick-tech/bilibot/pkg/services/dynamic"
	"github.com/Augenblick-tech/bilibot/pkg/task/basetask"
)

type BiliTask struct {
	basetask.BaseTask
	name      string
	spec      string
	BotID     string
	Mid       string
	lastPubTS uint64
}

func New(spec, mid, botID string) *BiliTask {
	return &BiliTask{
		name:      mid,
		spec:      spec,
		Mid:       mid,
		BotID:     botID,
		lastPubTS: 0,
	}
}

func NewWithAttr(spec string, attr map[string]interface{}) *BiliTask {
	return New(spec, attr["Mid"].(string), attr["BotID"].(string))
}

func (b *BiliTask) Run() {
	defer func() {
		if r := recover(); r != nil {
			if b.Status == basetask.Running {
				email.SendEmail(1, fmt.Sprintf("任务%s: 发生错误", b.name), r)
				b.Status = basetask.Warning
				panic(r)
			}
		}
	}()

	data, err := bilibot.GetDynamic(b.Mid, "")
	if err != nil {
		panic(err)
	}

	if b.lastPubTS == 0 {
		dynm, err := dynamic.GetByMid(b.Mid, 1)
		if err != nil {
			panic(err)
		}
		if len(dynm) == 0 {
			b.lastPubTS = 1
		} else {
			b.lastPubTS = dynm[0].PubTS
		}
	}

	if len(data) > 0 && data[0].Modules.Author.PubTS > b.lastPubTS {
		log.Println("新动态", data[0].Modules.Content.Desc.Text)
		convetStr, err := plugin.UnicodeToStr(data[0].Modules.Content.Desc.Text)
		if err != nil {
			panic(err)
		}

		if convetStr != "" {
			Bot, err := bot.Get(b.BotID)
			if err != nil {
				panic(err)
			}
			resp, err := bilibot.DynamicReply(Bot.Cookie, data[0].ID, convetStr)
			if err != nil {
				panic(err)
			}
			if resp.Code != 0 {
				panic(resp)
			}
		}

		email.SendEmail(1, "有新的动态！", fmt.Sprintf("%s:\n%s", data[0].Modules.Author.Name, data[0].Modules.Content.Desc.Text))
		b.lastPubTS = data[0].Modules.Author.PubTS
	}

	panic(dynamic.Add(data...))
}

func (b *BiliTask) Name() string {
	return b.name
}

func (b *BiliTask) Data() interface{} {
	return b.lastPubTS
}

func (b *BiliTask) Attribute() interface{} {
	return struct {
		Mid   string
		BotID string
	}{
		Mid:   b.Mid,
		BotID: b.BotID,
	}
}

func (b *BiliTask) SetStatus(s basetask.Status) {
	b.Status = s
}

func (b *BiliTask) GetStatus() basetask.Status {
	return b.Status
}

func (b *BiliTask) Spec() string {
	return b.spec
}
