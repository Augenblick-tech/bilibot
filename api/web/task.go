package web

import (
	"strconv"

	"github.com/Augenblick-tech/bilibot/lib/engine"
	"github.com/Augenblick-tech/bilibot/lib/task"
	"github.com/Augenblick-tech/bilibot/pkg/e"
	"github.com/Augenblick-tech/bilibot/pkg/task/basetask"
)

func NewTask(c *engine.Context) (interface{}, error) {
	return nil, nil
}

// SetTaskStatus godoc
// @Summary     更新任务状态
// @Description
// @Tags        web
// @Accept      json
// @Produce     json
// @Security 	ApiKeyAuth
// @Param       task		query	string	true	"任务名称"
// @Param       status		query	string	true	"任务状态"
// @Router      /web/task/update [post]
func SetTaskStatus(c *engine.Context) (interface{}, error) {
	taskName := c.Query("task")
	taskStatus := c.Query("status")

	statusID, err := strconv.Atoi(taskStatus)
	if err != nil {
		return nil, err
	}

	tw := task.Task(taskName)
	if tw == nil {
		return nil, e.ErrInvalidParam
	}

	tw.Task().SetStatus(basetask.Status(statusID))

	return nil, nil
}

// GetTaskStatus godoc
// @Summary     获取任务状态
// @Description
// @Tags        web
// @Accept      json
// @Produce     json
// @Security 	ApiKeyAuth
// @Param       task		query	string	true	"任务名称"
// @Router      /web/task/status [get]
func GetTaskStatus(c *engine.Context) (interface{}, error) {
	taskName := c.Query("task")

	tw := task.Task(taskName)
	if tw == nil {
		return nil, e.ErrInvalidParam
	}

	return tw.Task().GetStatus().String(), nil
}
