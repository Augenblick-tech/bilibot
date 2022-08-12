package bilitask

import (
	"context"
	"time"

	"github.com/Augenblick-tech/bilibot/lib/bili_bot"
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
			Mid:        mid,
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
			temp, err := bilibot.GetDynamic(b.Mid)
			if err != nil {
				b.E = err
				b.TaskStatus = tasks.TaskStatus_Error
				return
			}
			b.data = temp
		}
	}
}

func (b *biliTask) Stop() error {
	if b.TaskStatus == tasks.TaskStatus_Running {
		b.cancel()
	}
	return nil
}
