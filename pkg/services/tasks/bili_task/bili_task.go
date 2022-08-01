package bilitask

import (
	"context"
	"time"

	"github.com/Augenblick-tech/bilibot/pkg/model"
	"github.com/Augenblick-tech/bilibot/pkg/services/process"
)

type biliTask struct {
	*process.BaseTask
	data   []model.Dynamic
	ctx    context.Context
	cancel context.CancelFunc
	ticker *time.Ticker
}

func NewBiliTask(mid string, d time.Duration) *biliTask {
	return &biliTask{
		ticker: time.NewTicker(d),
		data:   make([]model.Dynamic, 0),
		BaseTask: &process.BaseTask{
			TaskStatus: process.TaskStatus_NotRunning,
			Mid:        mid,
		},
	}
}

func (b *biliTask) Data() interface{} {
	return b.data
}

func (b *biliTask) Run() {
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

func (b *biliTask) Stop() error {
	if b.TaskStatus == process.TaskStatus_Running {
		b.cancel()
	}
	return nil
}
