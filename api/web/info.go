package web

import (
	"github.com/Augenblick-tech/bilibot/lib/engine"
	"github.com/Augenblick-tech/bilibot/pkg/services/author"
	"github.com/Augenblick-tech/bilibot/pkg/services/bot"
	"github.com/Augenblick-tech/bilibot/pkg/services/dynamic"
	"github.com/Augenblick-tech/bilibot/pkg/services/user"
)

// GetBotList godoc
// @Summary      获取 Bot 列表
// @Description  根据 Token 获取 Bot 列表
// @Tags         web
// @Produce      json
// @Param 		 Authorization 	header 	string	true "Bearer 用户令牌"
// @Router       /web/bot/list [get]
func GetBotList(c *engine.Context) (interface{}, error) {
	id := c.Context.GetUint("UserID")
	return bot.GetList(id)
}

// GetAuthorList godoc
// @Summary      获取 up 主列表
// @Description
// @Tags         web
// @Produce      json
// @Param 		 Authorization 	header 	string	true	"Bearer 用户令牌"
// @Param        bot_id			query	string	true	"BotID"
// @Router       /web/author/list [get]
func GetAuthorList(c *engine.Context) (interface{}, error) {
	id := c.Context.GetUint("UserID")
	BotID := c.Query("bot_id")

	if err := user.CheckRecordWithID(id, BotID); err != nil {
		return nil, err
	}

	return author.GetList(BotID)
}

// GetDynamicList godoc
// @Summary      获取 up 主的动态列表
// @Description
// @Tags         web
// @Produce      json
// @Param 		 Authorization 	header 	string	true	"Bearer 用户令牌"
// @Param        bot_id			query	string	true	"BotID"
// @Param        mid			query	string	true	"up主id"
// @Router       /web/dynamic/list [get]
func GetDynamicList(c *engine.Context) (interface{}, error) {
	id := c.Context.GetUint("UserID")
	AuthorID := c.Query("mid")
	BotID := c.Query("bot_id")

	if err := user.CheckRecordWithID(id, BotID, AuthorID); err != nil {
		return nil, err
	}

	return dynamic.GetList(AuthorID)
}



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
// @Router       /web/author/add [post]
func AddAuthor(c *engine.Context) (interface{}, error) {
	id := c.Context.GetUint("UserID")
	info := addAuthorInfo{}

	if err := c.Bind(&info); err != nil {
		return nil, err
	}

	if err := user.CheckRecordWithID(id, info.BotID, info.Mid); err != nil {
		return nil, err
	}

	return nil, author.Add(info.Mid, info.BotID)
}
