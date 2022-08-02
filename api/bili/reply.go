package bili

import (
	"github.com/Augenblick-tech/bilibot/lib/engine"
	"github.com/Augenblick-tech/bilibot/pkg/model"
)

func AddReply(c *engine.Context) (interface{}, error) {
	var comment = struct {
		Type    string
		Oid     string
		Message string
	}{}

	c.Bind(&comment)

	replyResp, err := model.AddReply(comment.Type, comment.Oid, comment.Message)
	if err != nil {
		return nil, err
	}

	return replyResp, nil
}
