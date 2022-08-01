package process

import (
	"fmt"
	"sort"
	"sync"
)

type process struct {
	sync.RWMutex
	tasks map[string]*pTask
	start chan struct{}
}

type pTask struct {
	index int
	task  Task
}

type DelError []Task

func (e DelError) Error() string {
	s := ""
	d := ","
	for n, i := range e {
		t := i.(Task)
		if n == len(e)-1 {
			d = ""
		}
		s = fmt.Sprintf("%s%s%s", s, t.Error(), d)
	}
	s = fmt.Sprintf("[%s]", s)
	return s
}

type Task interface {
	// key
	SetMid(string)
	GetMid() string

	Status() TaskStatus
	GetError() error
	SetError(error)
	Error() string

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
		tasks: make(map[string]*pTask),
		start: make(chan struct{}),
	}
	go Process.run()
}

// 添加任务
func (p *process) Add(ts ...Task) {
	p.Lock()
	defer p.Unlock()
	for _, t := range ts {
		if _, ok := p.tasks[t.GetMid()]; ok {
			continue
		}
		p.tasks[t.GetMid()] = &pTask{
			index: len(p.tasks) + 1,
			task:  t,
		}
	}
	p.start <- struct{}{}
}

// 删除任务
func (p *process) Dels(mids ...string) error {
	p.Lock()
	defer p.Unlock()
	e := make(DelError, 0)
	for _, mid := range mids {
		t, ok := p.tasks[mid]
		if !ok {
			continue
		}
		err := t.task.Stop()
		if err != nil {
			t.task.SetError(err)
			e = append(e, t.task)
			continue
		}
		delete(p.tasks, t.task.GetMid())
	}

	if len(e) > 0 {
		return e
	}

	return nil
}

func (p *process) Del(mid string) error {
	p.Lock()
	defer p.Unlock()
	t, ok := p.tasks[mid]
	if !ok {
		return nil
	}
	err := t.task.Stop()
	if err != nil {
		t.task.SetError(err)
		return t.task
	}
	delete(p.tasks, t.task.GetMid())
	return nil
}

func (p *process) Stop(mid string) error {
	p.Lock()
	defer p.Unlock()

	if t, ok := p.tasks[mid]; ok {
		return t.task.Stop()
	}

	return nil
}

func (p *process) Status(mids ...string) []Task {
	p.RLock()
	defer p.RUnlock()
	l := make([]*pTask, 0)
	if len(mids) <= 0 {
		for _, s := range p.tasks {
			l = append(l, s)
		}
	} else {
		for _, id := range mids {
			s, ok := p.tasks[id]
			if !ok {
				continue
			}
			l = append(l, s)
		}
	}

	if len(l) > 1 {
		sort.Slice(l, func(i, j int) bool {
			return l[i].index > l[j].index
		})
	}

	r := make([]Task, 0, len(l))
	for _, t := range l {
		r = append(r, t.task)
	}

	return r
}

func (p *process) Run(mid string) error {
	t, ok := p.tasks[mid]
	if !ok {
		return fmt.Errorf("%s 不存在", mid)
	}

	if t.task.Status() != TaskStatus_Running {
		go t.task.Run()
	}

	return nil
}

func (p *process) run() {
	defer func() {
		if r := recover(); r != nil {
			// log
		}
		p.run()
	}()

	for {
		select {
		case <-p.start:
			{
				for _, t := range p.tasks {
					if t.task.Status() != TaskStatus_NotRunning {
						go t.task.Run()
					}
				}
			}
		}
	}
}

type BaseTask struct {
	TaskStatus TaskStatus
	E          error
	Mid        string
}

func (b *BaseTask) SetMid(mid string) {
	b.Mid = mid
}

func (b *BaseTask) GetMid() string {
	return b.Mid
}

func (b *BaseTask) Status() TaskStatus {
	return b.TaskStatus
}

func (b *BaseTask) SetError(e error) {
	b.E = e
}

func (b *BaseTask) GetError() error {
	return b.E
}

func (b *BaseTask) Error() string {
	e := ""
	if b.E != nil {
		e = b.E.Error()
	}
	return fmt.Sprintf("{\"mid\": \"%s\", \"error\": \"%s\"}", b.Mid, e)
}
