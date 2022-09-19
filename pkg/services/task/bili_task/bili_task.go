package bilitask

import (
	"log"

	"github.com/Augenblick-tech/bilibot/lib/bili_bot"
	"github.com/Augenblick-tech/bilibot/pkg/email"
	"github.com/Augenblick-tech/bilibot/pkg/services/dynamic"
)

type BiliTask struct {
	name      string
	spec      string
	Mid       string
	lastPubTS uint64
}

func New(spec, mid string) *BiliTask {
	return &BiliTask{
		name:      mid,
		spec:      spec,
		Mid:       mid,
		lastPubTS: 0,
	}
}

func NewWithAttr(spec string, attr map[string]interface{}) *BiliTask {
	return New(spec, attr["Mid"].(string))
}

func (b *BiliTask) Run() {
	data, err := bilibot.GetDynamic(b.Mid, "")
	if err != nil {
		panic(err)
	}
	
	if b.lastPubTS == 0 {
		dynm, err := dynamic.GetByMid(b.Mid, 1)
		if err != nil {
			panic(err)
		}
		b.lastPubTS = dynm[0].PubTS
	}

	if data[0].Modules.Author.PubTS > b.lastPubTS {
		log.Println("新动态", data[0].Modules.Content.Desc.Text)
		email.SendEmail(1, "有新的动态！", data[0].Modules.Content.Desc.Text)
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
		Mid string
	}{
		Mid: b.Mid,
	}
}

func (b *BiliTask) Spec() string {
	return b.spec
}
