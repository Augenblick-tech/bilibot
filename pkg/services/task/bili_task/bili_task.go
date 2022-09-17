package bilitask

import (
	"log"

	"github.com/Augenblick-tech/bilibot/lib/bili_bot"
	"github.com/Augenblick-tech/bilibot/pkg/services/dynamic"
)

type biliTask struct {
	name string
	spec string
	mid  string
	data []bilibot.Dynamic
}

func New(spec, mid string) *biliTask {
	return &biliTask{
		name: mid,
		spec: spec,
		mid:  mid,
		data: make([]bilibot.Dynamic, 0),
	}
}

func (b *biliTask) Run() {
	log.Println("Run biliTask")
	data, err := bilibot.GetDynamic(b.mid, "")
	if err != nil {
		panic(err)
	}
	log.Println(data)
	if len(b.data) == 0 {
		b.data = data
	} else {
		// 初步剪切动态(仅限运行时)
		for i, v := range data {
			// 从data中选取最新动态与temp中最新动态开始比较
			if v.Modules.Author.PubTS == b.data[0].Modules.Author.PubTS {
				// 避免空数据
				if i > 0 {
					panic(dynamic.Add(data[:i]...))
				}
				break // 截断最新动态后结束循环
			}
		}
		// 捡漏，如果temp最老的动态发布时间都大于data的最老发布时间就全部添加
		if data[len(data)-1].Modules.Author.PubTS > b.data[0].Modules.Author.PubTS {
			panic(dynamic.Add(data...))
		}
	}
}

func (b *biliTask) Name() string {
	return b.name
}

func (b *biliTask) Data() interface{} {
	return b.data
}

func (b *biliTask) Spec() string {
	return b.spec
}
