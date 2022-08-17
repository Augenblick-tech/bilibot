package bili

import (
	"github.com/Augenblick-tech/bilibot/lib/bili_bot"
	"github.com/Augenblick-tech/bilibot/lib/engine"
	"github.com/Augenblick-tech/bilibot/pkg/services/bot"
)

type commentInfo struct {
	BotID   string `json:"bot_id" binding:"required"`
	Type    int    `json:"type" binding:"required"`
	Oid     string `json:"oid" binding:"required"`
	Message string `json:"message" binding:"required"`
}

// AddReply godoc
// @Summary      根据type与oid进行回复
// @Description
// @Tags         bili
// @Accept       json
// @Produce      json
// @Param 		 Authorization 	header 	string			true	"Bearer 用户令牌"
// @Param        object			body	commentInfo		true	"回复评论详细信息"
// @Router       /bili/reply/add [post]
func AddReply(c *engine.Context) (interface{}, error) {
	var comment = commentInfo{}

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
