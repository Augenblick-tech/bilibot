package bili

import (
	"github.com/Augenblick-tech/bilibot/lib/bili_bot"
	"github.com/Augenblick-tech/bilibot/lib/engine"
)

// GetDynamic godoc
// @Summary     获取动态列表(访问b站api)
// @Description
// @Tags        bili
// @Produce     json
// @Security 	ApiKeyAuth
// @Param       mid			query		string	true	"up主id"
// @Param       offset		query		string	false	"动态偏移量"
// @Success		200			{array}		bilibot.Dynamic
// @Router      /bili/dynamic/getDynamic [get]
func GetDynamic(c *engine.Context) (interface{}, error) {
	AuthorID := c.Query("mid")
	Offset := c.Query("offset")

	return bilibot.GetDynamic(AuthorID, Offset)
}
