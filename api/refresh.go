package api

import (
	"time"

	"github.com/Augenblick-tech/bilibot/lib/engine"
	"github.com/Augenblick-tech/bilibot/pkg/e"
	"github.com/Augenblick-tech/bilibot/pkg/services/tasks/bili_task"
	"github.com/spf13/viper"
)

var biliTasks = make(map[string]*bilitask.BiliTask)

func RefreshDynamic(c *engine.Context) (interface{}, error) {

	mid := c.Query("mid")

	if _, ok := biliTasks[mid]; ok {
		return nil, e.RespCode_AlreadyExist
	}

	biliTask := bilitask.NewBiliTask(mid, time.Second*time.Duration(viper.GetInt("user.RefreshTime")))
	go biliTask.Run()
	
	biliTasks[mid] = biliTask

	return "success", nil
}

func GetLatestDynamic(c *engine.Context) (interface{}, error) {
	mid := c.Query("mid")

	if biliTask, ok := biliTasks[mid]; ok {
		return biliTask.Data(), nil
	}

	return nil, e.RespCode_ParamError
}

func GetStatus(c *engine.Context) (interface{}, error) {
	mid := c.Query("mid")

	if biliTask, ok := biliTasks[mid]; ok {
		return biliTask.TaskStatus, nil
	} else {
		return nil, e.RespCode_ParamError
	}
}

func StopRefreshDynamic(c *engine.Context) (interface{}, error) {
	mid := c.Query("mid")

	if biliTask, ok := biliTasks[mid]; ok {
		err := biliTask.Stop()
		if err != nil {
			return nil, err
		}
		return "success", nil
	} else {
		return nil, e.RespCode_ParamError
	}
}
