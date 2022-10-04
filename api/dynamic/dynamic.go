package dynamic

import (
	"fmt"

	"github.com/Augenblick-tech/bilibot/lib/engine"
	"github.com/Augenblick-tech/bilibot/lib/task"
	"github.com/Augenblick-tech/bilibot/pkg/e"
	"github.com/Augenblick-tech/bilibot/pkg/services/dynamic"
	"github.com/Augenblick-tech/bilibot/pkg/services/user"
	"github.com/Augenblick-tech/bilibot/pkg/task/bili/bilitask"
	"github.com/Augenblick-tech/bilibot/pkg/task/bili/check"
)

// Listen godoc
// @Summary     监听up主动态
// @Description 根据设定的时间间隔监听up主动态
// @Tags        web
// @Produce     json
// @Security 	ApiKeyAuth
// @Param       bot_id		query	string	true	"BotID"
// @Param       mid			query	string	true	"up主ID"
// @Param       interval	query	string	true	"监听时间间隔(s)"
// @Router      /web/dynamic/listen [get]
func Listen(c *engine.Context) (interface{}, error) {
	id := c.Context.GetUint("UserID")
	BotID := c.Query("bot_id")
	Mid := c.Query("mid")
	interval := c.Query("interval")

	if err := user.CheckRecordWithID(id, BotID, Mid); err != nil {
		return nil, err
	}

	b := bilitask.New(fmt.Sprintf("@every %ss", interval), Mid, BotID)
	checker := check.New("@every 30s", BotID)

	return task.Add(c.Context.GetUint("UserID"), b, checker)
}

// Latest godoc
// @Summary     获取up主最新动态
// @Description
// @Tags        web
// @Produce     json
// @Security 	ApiKeyAuth
// @Param       bot_id		query	string	true	"BotID"
// @Param       mid			query	string	true	"up主ID"
// @Router      /web/dynamic/latest [get]
func Latest(c *engine.Context) (interface{}, error) {
	id := c.Context.GetUint("UserID")
	BotID := c.Query("bot_id")
	Mid := c.Query("mid")

	if err := user.CheckRecordWithID(id, BotID, Mid); err != nil {
		return nil, err
	}

	dynm, err := dynamic.GetByMid(Mid, 1)
	return dynm[0], err
}

// Status godoc
// @Summary     获取传入的uid的状态
// @Description
// @Tags        web
// @Produce     json
// @Security 	ApiKeyAuth
// @Param       bot_id		query	string	true	"BotID"
// @Param       mid			query	string	true	"up主ID"
// @Router      /web/dynamic/status [get]
func Status(c *engine.Context) (interface{}, error) {
	id := c.Context.GetUint("UserID")
	BotID := c.Query("bot_id")
	Mid := c.Query("mid")

	if err := user.CheckRecordWithID(id, BotID, Mid); err != nil {
		return nil, err
	}

	dynm := task.Task(Mid)
	if dynm == nil {
		return nil, e.ErrNotFound
	}
	return dynm, nil
}

// Stop godoc
// @Summary     停止传入的uid的任务
// @Description
// @Tags        web
// @Produce     json
// @Security 	ApiKeyAuth
// @Param       bot_id		query	string	true	"BotID"
// @Param       mid			query	string	true	"up主ID"
// @Router      /web/dynamic/stop [get]
func Stop(c *engine.Context) (r interface{}, err error) {
	id := c.Context.GetUint("UserID")
	BotID := c.Query("bot_id")
	Mid := c.Query("mid")

	if err := user.CheckRecordWithID(id, BotID, Mid); err != nil {
		return nil, err
	}

	task.Remove(Mid)
	task.Remove(BotID)
	return
}
