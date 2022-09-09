package scheduletask

import (
	"github.com/Augenblick-tech/bilibot/pkg/services/tasks"
	"github.com/robfig/cron/v3"
)

type ScheduleTask struct {
	*tasks.BaseTask
	Cron    cron.Cron
	Actions []Action
}

type Action struct {
	spec string
	cmd  func()
}

type EntryID int

func NewScheduleTask() *ScheduleTask {
	return &ScheduleTask{
		Cron: *cron.New(),
		BaseTask: &tasks.BaseTask{
			TaskStatus: tasks.TaskStatus_NotRunning,
			Id:         "",
			E:          nil,
		},
		Actions: make([]Action, 0),
	}
}

func (s *ScheduleTask) AddFunc(spec string, cmd func()) (EntryID, error) {
	s.Actions = append(s.Actions, Action{
		spec,
		cmd,
	})
	id, err := s.Cron.AddFunc(spec, cmd)
	return EntryID(id), err
}

func (s *ScheduleTask) Run() {
	s.Cron.Start()
}

func (s *ScheduleTask) Stop() {
	s.Cron.Stop()
}
