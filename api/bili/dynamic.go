package bili

import (
	"github.com/Augenblick-tech/bilibot/lib/engine"
	"github.com/Augenblick-tech/bilibot/lib/bili_bot"
)

func GetDynamic(c *engine.Context) (interface{}, error) {
	dynamics, err := bilibot.GetDynamic(c.Query("mid"))
	if err != nil {
		return nil, err
	}

	return dynamics, nil
}
