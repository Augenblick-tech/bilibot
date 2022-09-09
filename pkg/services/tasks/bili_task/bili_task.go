package bilitask

import (
	"context"
	"time"

	"github.com/Augenblick-tech/bilibot/lib/bili_bot"
	"github.com/Augenblick-tech/bilibot/pkg/services/dynamic"
	"github.com/Augenblick-tech/bilibot/pkg/services/tasks"
)

type biliTask struct {
	*tasks.BaseTask
	data   []bilibot.Dynamic
	ctx    context.Context
	cancel context.CancelFunc
	ticker *time.Ticker
}

func NewBiliTask(mid string, d time.Duration) *biliTask {
	return &biliTask{
		ticker: time.NewTicker(d),
		data:   make([]bilibot.Dynamic, 0),
		BaseTask: &tasks.BaseTask{
			TaskStatus: tasks.TaskStatus_NotRunning,
			Id:         mid,
			E:          nil,
		},
	}
}

func (b *biliTask) Data() interface{} {
	return b.data
}

func (b *biliTask) Run() {
	b.ctx, b.cancel = context.WithCancel(context.Background())
	b.TaskStatus = tasks.TaskStatus_Running
	for {
		select {
		case <-b.ctx.Done():
			b.TaskStatus = tasks.TaskStatus_Stoped
			return
		case <-b.ticker.C:
			temp, err := bilibot.GetDynamic(b.Id, "0")
			if err != nil {
				b.E = err
				b.TaskStatus = tasks.TaskStatus_Error
				return
			}
			// 初步剪切动态(仅限运行时)
			if len(b.data) > 0 {
				for i, v := range temp {
					// 从data中选取最新动态与temp中最新动态开始比较
					if v.Modules.Author.PubTS == b.data[0].Modules.Author.PubTS {
						// 避免空数据
						if i > 0 {
							b.SetError(dynamic.Add(temp[:i]...))
						}
						break // 截断最新动态后结束循环
					}
				}
				// 捡漏，如果temp最老的动态发布时间都大于data的最老发布时间就全部添加
				if temp[len(temp)-1].Modules.Author.PubTS > b.data[0].Modules.Author.PubTS {
					b.SetError(dynamic.Add(temp...))
				}
			} else {
				b.SetError(dynamic.Add(temp...))
			}
			b.data = temp
		}
	}
}

func (b *biliTask) Stop() error {
	if b.TaskStatus == tasks.TaskStatus_Running {
		b.cancel()
		return nil
	}
	return b.Status()
}
