package api

import (
	"time"

	"github.com/Augenblick-tech/bilibot/lib/engine"
	"github.com/Augenblick-tech/bilibot/pkg/e"
	"github.com/Augenblick-tech/bilibot/pkg/services/tasks/bili_task"
	"github.com/spf13/viper"
)

var biliTasks = make(map[string]*bilitask.BiliTask)

// RefreshDynamic godoc
// @Summary      监听up主动态
// @Description  根据设定的时间间隔监听up主动态
// @Tags         v2
// @Produce      json
// @Param        mid   query     string  true  "up主id"
// @Router       /dynamic/refresh [get]
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

// GetLatestDynamic godoc
// @Summary      获取up主最新动态
// @Description  
// @Tags         v2
// @Produce      json
// @Param        mid   query     string  true  "up主id"
// @Router       /dynamic/latest [get]
func GetLatestDynamic(c *engine.Context) (interface{}, error) {
	mid := c.Query("mid")

	if biliTask, ok := biliTasks[mid]; ok {
		return biliTask.Data(), nil
	}

	return nil, e.RespCode_ParamError
}

// GetStatus godoc
// @Summary      获取传入的uid的状态
// @Description  
// @Tags         v2
// @Produce      json
// @Param        mid   query     string  true  "up主id"
// @Router       /dynamic/status [get]
func GetStatus(c *engine.Context) (interface{}, error) {
	mid := c.Query("mid")

	if biliTask, ok := biliTasks[mid]; ok {
		return biliTask.TaskStatus, nil
	} else {
		return nil, e.RespCode_ParamError
	}
}

// StopRefreshDynamic godoc
// @Summary      停止传入的uid的任务
// @Description  
// @Tags         v2
// @Produce      json
// @Param        mid   query     string  true  "up主id"
// @Router       /dynamic/stop [get]
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
