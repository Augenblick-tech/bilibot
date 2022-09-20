package web

import (
	"strconv"

	"github.com/Augenblick-tech/bilibot/lib/engine"
	"github.com/Augenblick-tech/bilibot/lib/task"
	"github.com/Augenblick-tech/bilibot/pkg/task/basetask"
)

func NewTask(c *engine.Context) (interface{}, error) {
	return nil, nil
}

// SetTaskTatus godoc
// @Summary     更新任务状态
// @Description
// @Tags        web
// @Accept      json
// @Produce     json
// @Security 	ApiKeyAuth
// @Param       task		query	string	true	"任务名称"
// @Param       status		query	string	true	"任务状态"
// @Router      /web/task/update [post]
func SetTaskTatus(c *engine.Context) (interface{}, error) {
	taskName := c.Query("task")
	taskStatus := c.Query("status")

	statusID, err := strconv.Atoi(taskStatus)
	if err != nil {
		return nil, err
	}

	tw := task.Task(taskName)
	tw.Task().SetStatus(basetask.Status(statusID))

	return nil, nil
}
