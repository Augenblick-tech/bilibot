package dynamic

import (
	"time"

	"github.com/Augenblick-tech/bilibot/lib/engine"
	"github.com/Augenblick-tech/bilibot/pkg/e"
	"github.com/Augenblick-tech/bilibot/pkg/services/tasks"
	bilitask "github.com/Augenblick-tech/bilibot/pkg/services/tasks/bili_task"
	"github.com/spf13/viper"
)

// Listen godoc
// @Summary      监听up主动态
// @Description  根据设定的时间间隔监听up主动态
// @Tags         v2
// @Produce      json
// @Param        mid   query     string  true  "up主id"
// @Router       /dynamic/listen [get]
func Listen(c *engine.Context) (interface{}, error) {
	mid := c.Query("mid")

	if tasks.Process.IsExists(mid) {
		return nil, e.RespCode_AlreadyExist
	}

	tasks.Process.Add(
		bilitask.NewBiliTask(
			mid,
			time.Second*time.Duration(viper.GetInt("user.RefreshTime"))),
	)

	return nil, nil
}

// Latest godoc
// @Summary      获取up主最新动态
// @Description
// @Tags         v2
// @Produce      json
// @Param        mid   query     string  true  "up主id"
// @Router       /dynamic/latest [get]
func Latest(c *engine.Context) (interface{}, error) {
	status := tasks.Process.Status(c.Query("mid"))
	if len(status) <= 0 {
		return nil, e.RespCode_ParamError
	}
	return status[0].Data(), nil
}

// Status godoc
// @Summary      获取传入的uid的状态
// @Description
// @Tags         v2
// @Produce      json
// @Param        mid   query     string  true  "up主id"
// @Router       /dynamic/status [get]
func Status(c *engine.Context) (interface{}, error) {
	mid := c.Query("mid")

	status := tasks.Process.Status(mid)
	if len(status) > 0 {
		return status[0].Status(), nil
	}

	return nil, e.RespCode_ParamError

}

// Stop godoc
// @Summary      停止传入的uid的任务
// @Description
// @Tags         v2
// @Produce      json
// @Param        mid   query     string  true  "up主id"
// @Router       /dynamic/stop [get]
func Stop(c *engine.Context) (r interface{}, err error) {
	err = tasks.Process.Stop(c.Query("mid"))
	if err != nil {
		// print log
		return nil, e.RespCode_ParamError
	}
	return
}
