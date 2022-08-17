package dynamic

import (
	"time"

	"github.com/Augenblick-tech/bilibot/lib/engine"
	"github.com/Augenblick-tech/bilibot/pkg/e"
	"github.com/Augenblick-tech/bilibot/pkg/services/author"
	"github.com/Augenblick-tech/bilibot/pkg/services/tasks"
	bilitask "github.com/Augenblick-tech/bilibot/pkg/services/tasks/bili_task"
	"github.com/Augenblick-tech/bilibot/pkg/services/user"
	"github.com/spf13/viper"
)

type addAuthorInfo struct {
	Mid   string `json:"mid"`
	BotID string `json:"bot_id"`
}

// AddAuthor godoc
// @Summary      添加up主
// @Description  需先添加up主之后才能监听动态
// @Tags         web
// @Accept       json
// @Produce      json
// @Param 		 Authorization 	header 	string			true	"Bearer 用户令牌"
// @Param        object			body	addAuthorInfo	true	"up主id和BotID"
// @Router       /web/dynamic/addAuthor [post]
func AddAuthor(c *engine.Context) (interface{}, error) {
	id := c.Context.GetUint("UserID")
	info := addAuthorInfo{}

	if err := user.CheckRecordWithID(id, info.BotID, info.Mid); err != nil {
		return nil, err
	}

	return nil, author.Add(info.Mid, info.BotID)
}

// Listen godoc
// @Summary      监听up主动态
// @Description  根据设定的时间间隔监听up主动态
// @Tags         web
// @Produce      json
// @Param 		 Authorization 	header 	string			true	"Bearer 用户令牌"
// @Param        object			body	addAuthorInfo	true	"up主id和BotID"
// @Router       /web/dynamic/listen [get]
func Listen(c *engine.Context) (interface{}, error) {
	id := c.Context.GetUint("UserID")
	mid := c.Query("mid")
	botID := c.Query("bot_id")

	if err := user.CheckRecordWithID(id, botID, mid); err != nil {
		return nil, err
	}

	if !tasks.Process.IsExists(mid) {
		tasks.Process.Add(
			bilitask.NewBiliTask(
				mid,
				time.Second*time.Duration(viper.GetInt("user.RefreshTime"))),
		)
		return nil, nil
	} else {
		if tasks.Process.Status(mid)[0].Status() == tasks.TaskStatus_Stoped {
			return nil, tasks.Process.Run(mid)
		}
		return nil, tasks.Process.Status(mid)[0].Status()
	}
}

// Latest godoc
// @Summary      获取up主最新动态
// @Description
// @Tags         web
// @Produce      json
// @Param 		 Authorization 	header 	string			true	"Bearer 用户令牌"
// @Param        object			body	addAuthorInfo	true	"up主id和BotID"
// @Router       /web/dynamic/latest [get]
func Latest(c *engine.Context) (interface{}, error) {
	id := c.Context.GetUint("UserID")
	mid := c.Query("mid")
	botID := c.Query("bot_id")

	if err := user.CheckRecordWithID(id, botID, mid); err != nil {
		return nil, err
	}

	status := tasks.Process.Status(mid)
	if len(status) <= 0 {
		return nil, e.RespCode_ParamError
	}
	return status[0].Data(), nil
}

// Status godoc
// @Summary      获取传入的uid的状态
// @Description
// @Tags         web
// @Produce      json
// @Param 		 Authorization 	header 	string			true	"Bearer 用户令牌"
// @Param        object			body	addAuthorInfo	true	"up主id和BotID"
// @Router       /web/dynamic/status [get]
func Status(c *engine.Context) (interface{}, error) {
	id := c.Context.GetUint("UserID")
	mid := c.Query("mid")
	botID := c.Query("bot_id")

	if err := user.CheckRecordWithID(id, botID, mid); err != nil {
		return nil, err
	}

	status := tasks.Process.Status(mid)
	if len(status) > 0 {
		return status[0].Status(), nil
	}

	return nil, e.RespCode_ParamError

}

// Stop godoc
// @Summary      停止传入的uid的任务
// @Description
// @Tags         web
// @Produce      json
// @Param 		 Authorization 	header 	string			true	"Bearer 用户令牌"
// @Param        object			body	addAuthorInfo	true	"up主id和BotID"
// @Router       /web/dynamic/stop [get]
func Stop(c *engine.Context) (r interface{}, err error) {
	id := c.Context.GetUint("UserID")
	mid := c.Query("mid")
	botID := c.Query("bot_id")

	if err := user.CheckRecordWithID(id, botID, mid); err != nil {
		return nil, err
	}

	err = tasks.Process.Stop(mid)
	if err != nil {
		// print log
		return nil, err
	}
	return
}
