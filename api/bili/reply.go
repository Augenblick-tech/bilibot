package bili

import (
	bilibot "github.com/Augenblick-tech/bilibot/lib/bili_bot"
	"github.com/Augenblick-tech/bilibot/lib/engine"
	"github.com/Augenblick-tech/bilibot/pkg/model/api"
	"github.com/Augenblick-tech/bilibot/pkg/services/bot"
)

// AddReply godoc
// @Summary     根据type与oid进行回复
// @Description
// @Tags        bili
// @Accept      json
// @Produce     json
// @Security 	ApiKeyAuth
// @Param       object		body		api.CommentInfo	true	"回复评论详细信息"
// @Success		200			{object}	api.ReplyInfo
// @Router      /bili/reply/add [post]
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

	reply, err := bilibot.AddReply(Bot.Cookie, comment.Type, comment.Oid, comment.Message)
	if err != nil {
		return nil, err
	}

	return api.ReplyInfo{
		SuccessToast: reply.Data.SuccessToast,
		Emote:        reply.Data.Emote,
	}, nil
}
