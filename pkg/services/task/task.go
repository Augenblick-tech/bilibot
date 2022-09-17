package task

import (
	"context"
	"reflect"

	"github.com/robfig/cron/v3"
)

type Process struct {
	c     *cron.Cron
	tasks map[string]*TaskWrapper
}

type TaskWrapper struct {
	EntryID
	Name string
	t    *Job
	Type string
}

func (tw *TaskWrapper) Task() Job {
	return *tw.t
}

type EntryID int

var std = New(
	cron.WithSeconds(),
	cron.WithChain(cron.Recover(cron.DefaultLogger)),
)

var (
	disableDefaultProcess = false
)

func init() {
	if !disableDefaultProcess {
		std.Start()
	}
}

// func SetMode(mode string) {
// }

func New(opts ...cron.Option) *Process {
	return &Process{
		c:     cron.New(opts...),
		tasks: make(map[string]*TaskWrapper),
	}
}

func (P *Process) Start() {
	P.c.Start()
}

func (P *Process) Add(t Job) (EntryID, error) {
	id, err := P.c.AddFunc(t.Spec(), func() {
		t.Run()
	})
	if err != nil {
		return 0, err
	}

	entryID := EntryID(id)
	P.tasks[t.Name()] = &TaskWrapper{
		EntryID: entryID,
		Name: t.Name(),
		t:       &t,
		Type:    reflect.TypeOf(t).String(),
	}

	return entryID, err
}

func (P *Process) Tasks() []TaskWrapper {
	var tasks = []TaskWrapper{}
	for _, t := range P.tasks {
		tasks = append(tasks, *t)
	}

	return tasks
}

func (P *Process) Task(name string) *TaskWrapper {
	return P.tasks[name]
}

func (P *Process) Remove(name string) {
	P.c.Remove(cron.EntryID(P.tasks[name].EntryID))
	delete(P.tasks, name)
}

func (P *Process) Stop() context.Context {
	return P.c.Stop()
}

// std is the default process.
func Start() {
	std.Start()
}

func Add(t Job) (EntryID, error) {
	return std.Add(t)
}

func Tasks() []TaskWrapper {
	return std.Tasks()
}

func Task(name string) *TaskWrapper {
	return std.Task(name)
}

func Remove(name string) {
	std.Remove(name)
}

func Stop() context.Context {
	return std.Stop()
}

// Job is the interface that wraps the basic Run method.
type Job interface {
	cron.Job
	Name() string
	Data() interface{}
	Spec() string
}
