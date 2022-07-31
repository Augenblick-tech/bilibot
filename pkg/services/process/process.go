package process

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

type process struct {
	sync.RWMutex
	tasks map[string]Task
}

type Task interface {
	// 排序
	SetIndex(int)
	GetIndex() int

	// key
	SetMid(string)
	GetMid() string

	Status() TaskStatus
	Error() error

	// 执行的任务
	Run()

	// 停止
	Stop() error

	Data() interface{}
}

type TaskStatus int8

const (
	TaskStatus_Nil TaskStatus = iota
	TaskStatus_Running
	TaskStatus_NotRunning
	TaskStatus_Error
	TaskStatus_Stoped
)

var Process *process

func init() {
	Process = &process{
		tasks: make(map[string]Task),
	}
	go Process.Run()
}

// 添加任务
func (p *process) Add(t Task) error {
	p.Lock()
	defer p.Unlock()
	if _, ok := p.tasks[t.GetMid()]; ok {
		return fmt.Errorf("存在")
	}
	t.SetIndex(len(p.tasks) + 1)
	p.tasks[t.GetMid()] = t
	return nil
}

// 删除任务
func (p *process) Del(mid string) error {
	p.Lock()
	defer p.Unlock()

	t, ok := p.tasks[mid]
	if !ok {
		return fmt.Errorf("不存在")
	}
	err := t.Stop()
	if err != nil {
		return err
	}
	delete(p.tasks, t.GetMid())
	return nil
}

func (p *process) Stop(mid string) error {
	p.Lock()
	defer p.Unlock()

	if t, ok := p.tasks[mid]; ok {
		return t.Stop()
	}

	return nil
}

func (p *process) Status(mids ...string) []Task {
	p.RLock()
	defer p.RUnlock()
	r := make([]Task, 0)
	if len(mids) <= 0 {
		for _, s := range p.tasks {
			r = append(r, s)
		}
	} else {
		for _, id := range mids {
			s, ok := p.tasks[id]
			if !ok {
				continue
			}
			r = append(r, s)
		}
	}

	if len(r) > 1 {
		sort.Slice(r, func(i, j int) bool {
			return r[i].GetIndex() > r[j].GetIndex()
		})
	}

	return r
}

func (p *process) Run() {
	defer func() {
		if r := recover(); r != nil {
			// log
		}
		p.Run()
	}()

	for {
		time.Sleep(time.Second)
		p.Lock()
		for _, t := range p.tasks {
			if t.Status() != TaskStatus_NotRunning {
				go t.Run()
			}
		}
		p.Unlock()
	}
}
