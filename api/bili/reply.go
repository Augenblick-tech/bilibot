package bili

import (
	"github.com/Augenblick-tech/bilibot/lib/engine"
	"github.com/Augenblick-tech/bilibot/pkg/model"
)

func AddReply(c *engine.Context) (interface{}, error) {
	var comment = struct {
		Type    string `json:"type" binding:"required"`
		Oid     string `json:"oid" binding:"required"`
		Message string `json:"message" binding:"required"`
	}{}

	err := c.Bind(&comment)
	if err != nil {
		return nil, err
	}

	replyResp, err := model.AddReply(comment.Type, comment.Oid, comment.Message)
	if err != nil {
		return nil, err
	}

	return replyResp, nil
}
