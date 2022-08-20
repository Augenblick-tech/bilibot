package bili

import (
	"github.com/Augenblick-tech/bilibot/lib/bili_bot"
	"github.com/Augenblick-tech/bilibot/lib/engine"
	"github.com/Augenblick-tech/bilibot/pkg/model/api"
	"github.com/Augenblick-tech/bilibot/pkg/services/bot"
)

// AddReply godoc
// @Summary      根据type与oid进行回复
// @Description
// @Tags         bili
// @Accept       json
// @Produce      json
// @Param 		 Authorization 	header 	string			true	"Bearer 用户令牌"
// @Param        object			body	api.CommentInfo	true	"回复评论详细信息"
// @Router       /bili/reply/add [post]
func AddReply(c *engine.Context) (interface{}, error) {
	var comment = api.CommentInfo{}

	err := c.Bind(&comment)
	if err != nil {
		return nil, err
	}

	Bot, err := bot.Get(comment.BotID)
	if err != nil {
		return nil, err
	}

	return bilibot.AddReply(Bot.Cookie, comment.Type, comment.Oid, comment.Message)
}
