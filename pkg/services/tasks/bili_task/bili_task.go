package bilitask

import (
	"context"
	"time"

	"github.com/Augenblick-tech/bilibot/pkg/model"
	"github.com/Augenblick-tech/bilibot/pkg/services/process"
)

type BiliTask struct {
	*process.BaseTask
	data   []model.Dynamic
	ctx    context.Context
	cancel context.CancelFunc
	ticker *time.Ticker
}

func NewBiliTask(mid string, d time.Duration) *BiliTask {
	return &BiliTask{
		ticker: time.NewTicker(d),
		data:   make([]model.Dynamic, 0),
		BaseTask: &process.BaseTask{
			TaskStatus: process.TaskStatus_NotRunning,
			Mid:        mid,
		},
	}
}

func (b *BiliTask) Data() interface{} {
	return b.data
}

func (b *BiliTask) Run() {
	b.ctx, b.cancel = context.WithCancel(context.Background())
	b.TaskStatus = process.TaskStatus_Running
	for {
		select {
		case <-b.ctx.Done():
			b.TaskStatus = process.TaskStatus_Stoped
			return
		case <-b.ticker.C:
			temp, err := model.GetDynamic(b.Mid)
			if err != nil {
				b.E = err
				b.TaskStatus = process.TaskStatus_Error
				return
			}
			b.data = temp
		}
	}
}

func (b *BiliTask) Stop() error {
	if b.TaskStatus == process.TaskStatus_Running {
		b.cancel()
	}
	return nil
}
