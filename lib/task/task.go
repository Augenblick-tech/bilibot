package task

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"runtime"

	"github.com/Augenblick-tech/bilibot/pkg/dao"
	"github.com/Augenblick-tech/bilibot/pkg/model"
	"github.com/Augenblick-tech/bilibot/pkg/task/basetask"
	"github.com/Augenblick-tech/bilibot/pkg/task/bili/bilitask"
	"github.com/Augenblick-tech/bilibot/pkg/task/bili/check"
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
	cron.WithChain(Recover(cron.DefaultLogger), cron.DelayIfStillRunning(cron.DefaultLogger)),
)

func Recover(logger cron.Logger) cron.JobWrapper {
	return func(j cron.Job) cron.Job {
		return cron.FuncJob(func() {
			defer func() {
				if r := recover(); r != nil {
					const size = 64 << 10
					buf := make([]byte, size)
					buf = buf[:runtime.Stack(buf, false)]
					err, ok := r.(error)
					if !ok {
						err = fmt.Errorf("%v", r)
					}
					logger.Error(err, "panic", "stack", "...\n"+string(buf))

					// TODO: send email to admin
				}
			}()
			j.Run()
		})
	}
}

// var (
// 	disableDefaultProcess = false
// )

// func init() {
// 	if !disableDefaultProcess {
// 		std.Start()
// 	}
// }

// func SetMode(mode string) {
// }

func New(opts ...cron.Option) *Process {
	return &Process{
		c:     cron.New(opts...),
		tasks: make(map[string]*TaskWrapper),
	}
}

func (P *Process) Start() {
	tasks, err := dao.GetAllTask()
	if err != nil {
		panic(err)
	}

	attr := map[string]interface{}{}

	for _, t := range tasks {
		if t.Attribute != "" {
			if err := json.Unmarshal([]byte(t.Attribute), &attr); err != nil {
				panic(err)
			}
			switch t.Type {
			case "*bilitask.BiliTask":
				P.Add(t.UserID, bilitask.NewWithAttr(t.Spec, attr))
			case "*checklogin.BotLoginInfo":
				P.Add(t.UserID, check.NewWithAttr(t.Spec, attr))
			}
		}
	}

	P.c.Start()
}

func (P *Process) Add(UserID uint, tasks ...Job) (int, error) {
	taskCnt := 0

	for _, t := range tasks {
		if _, ok := P.tasks[t.Name()]; ok {
			continue
		}
		id, err := P.c.AddFunc(t.Spec(), func() {
			t.Run()
		})
		if err != nil {
			return 0, err
		}

		entryID := EntryID(id)
		tw := &TaskWrapper{
			EntryID: entryID,
			Name:    t.Name(),
			t:       &t,
			Type:    reflect.TypeOf(t).String(),
		}

		P.tasks[t.Name()] = tw

		attr, err := json.Marshal(t.Attribute())
		if err != nil {
			panic(err)
		}

		dao.CreateWithIgonreConflict(&model.Task{
			Name:      t.Name(),
			Spec:      t.Spec(),
			Type:      reflect.TypeOf(t).String(),
			Attribute: string(attr),
			UserID:    UserID,
		})
	}

	return taskCnt, nil
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
	if _, ok := P.tasks[name]; !ok {
		return
	}
	P.c.Remove(cron.EntryID(P.tasks[name].EntryID))
	delete(P.tasks, name)
	dao.Delete(&model.Task{
		Name: name,
	})
}

func (P *Process) Stop() context.Context {
	return P.c.Stop()
}

// std is the default process.
func Start() {
	std.Start()
}

func Add(UserID uint, tasks ...Job) (int, error) {
	return std.Add(UserID, tasks...)
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
	Name() string // unique field
	Attribute() interface{}
	SetStatus(basetask.Status)
	Data() interface{}
	Spec() string
}
