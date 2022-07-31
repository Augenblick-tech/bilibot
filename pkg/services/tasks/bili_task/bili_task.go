package bilitask

import (
	"context"
	"time"

	"github.com/lonzzi/bilibot/pkg/model"
	"github.com/lonzzi/bilibot/pkg/services/process"
)

type biliTask struct {
	index  int
	status process.TaskStatus
	e      error
	mid    string
	data   []model.Dynamic
	ctx    context.Context
	cancel context.CancelFunc
	ticker *time.Ticker
}

func NewBiliTask(mid string, d time.Duration) *biliTask {
	return &biliTask{
		mid:    mid,
		ticker: time.NewTicker(d),
		data:   make([]model.Dynamic, 0),
		status: process.TaskStatus_NotRunning,
	}
}

func (b *biliTask) SetIndex(i int) {
	b.index = i
}

func (b *biliTask) GetIndex() int {
	return b.index
}

func (b *biliTask) SetMid(mid string) {
	b.mid = mid
}

func (b *biliTask) GetMid() string {
	return b.mid
}

func (b *biliTask) Status() process.TaskStatus {
	return b.status
}

func (b *biliTask) SetError(e error) {
	b.e = e
}

func (b *biliTask) Error() error {
	return b.e
}

func (b *biliTask) Data() interface{} {
	return b.data
}

func (b *biliTask) Run() {
	b.ctx, b.cancel = context.WithCancel(context.Background())
	b.status = process.TaskStatus_Running
	for {
		select {
		case <-b.ctx.Done():
			b.status = process.TaskStatus_Stoped
			return
		case <-b.ticker.C:
			temp, err := model.GetDynamic(b.mid)
			if err != nil {
				b.e = err
				b.status = process.TaskStatus_Error
				return
			}
			b.data = temp
		}
	}
}

func (b *biliTask) Stop() error {
	if b.status == process.TaskStatus_Running {
		b.cancel()
	}
	return nil
}
