package bili

import (
	"github.com/Augenblick-tech/bilibot/lib/engine"
	"github.com/Augenblick-tech/bilibot/pkg/model"
)

func AddReply(c *engine.Context) (interface{}, error) {
	comment := c.PostBody()
	commentType := comment["type"].(string)
	oid := comment["oid"].(string)
	message := comment["message"].(string)
	replyResp, err := model.AddReply(commentType, oid, message)
	if err != nil {
		return nil, err
	}

	return replyResp, nil
}
