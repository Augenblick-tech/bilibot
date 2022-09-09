package tasks

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
		t := i
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
	SetId(string)
	GetId() string

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

//go:generate stringer -type=TaskStatus --linecomment
type TaskStatus int8

const (
	TaskStatus_Nil        TaskStatus = iota // 任务不存在
	TaskStatus_Running                      // 任务运行中
	TaskStatus_NotRunning                   // 任务未运行
	TaskStatus_Error                        // 任务出错
	TaskStatus_Stoped                       // 任务已停止
)

func (t TaskStatus) Error() string {
	return t.String()
}

var Process *process

func init() {
	Process = &process{
		tasks: make(map[string]*pTask),
		start: make(chan struct{}),
	}
	go Process.run()
}

func (p *process) IsExists(id string) bool {
	p.RLock()
	defer p.RUnlock()

	_, ok := p.tasks[id]
	return ok
}

// 添加任务
func (p *process) Add(ts ...Task) {
	p.Lock()
	defer p.Unlock()
	for _, t := range ts {
		if _, ok := p.tasks[t.GetId()]; ok {
			continue
		}
		p.tasks[t.GetId()] = &pTask{
			index: len(p.tasks) + 1,
			task:  t,
		}
	}
	p.start <- struct{}{}
}

// 删除任务
func (p *process) Dels(ids ...string) error {
	p.Lock()
	defer p.Unlock()
	e := make(DelError, 0)
	for _, id := range ids {
		t, ok := p.tasks[id]
		if !ok {
			continue
		}
		err := t.task.Stop()
		if err != nil {
			t.task.SetError(err)
			e = append(e, t.task)
			continue
		}
		delete(p.tasks, t.task.GetId())
	}

	if len(e) > 0 {
		return e
	}

	return nil
}

func (p *process) Del(id string) error {
	p.Lock()
	defer p.Unlock()
	t, ok := p.tasks[id]
	if !ok {
		return nil
	}
	err := t.task.Stop()
	if err != nil {
		t.task.SetError(err)
		return t.task
	}
	delete(p.tasks, t.task.GetId())
	return nil
}

func (p *process) Stop(id string) error {
	p.Lock()
	defer p.Unlock()

	if t, ok := p.tasks[id]; ok {
		return t.task.Stop()
	}

	return TaskStatus_Nil
}

func (p *process) Status(ids ...string) []Task {
	p.RLock()
	defer p.RUnlock()
	l := make([]*pTask, 0)
	if len(ids) <= 0 {
		for _, s := range p.tasks {
			l = append(l, s)
		}
	} else {
		for _, id := range ids {
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

func (p *process) Run(id string) error {
	t, ok := p.tasks[id]
	if !ok {
		return fmt.Errorf("%s 不存在", id)
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
					if t.task.Status() == TaskStatus_NotRunning {
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
	Id         string
}

func (b *BaseTask) SetId(id string) {
	b.Id = id
}

func (b *BaseTask) GetId() string {
	return b.Id
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
	return fmt.Sprintf("{\"mid\": \"%s\", \"error\": \"%s\"}", b.Id, e)
}
