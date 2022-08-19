package bili

import (
	"github.com/Augenblick-tech/bilibot/lib/engine"
	"github.com/Augenblick-tech/bilibot/lib/bili_bot"
)

// GetDynamic godoc
// @Summary      获取动态列表(访问b站api)
// @Description
// @Tags         bili
// @Produce      json
// @Param 		 Authorization 	header 	string	true	"Bearer 用户令牌"
// @Param        mid			query	string	true	"up主id"
// @Router       /bili/dynamic/getDynamic [get]
func GetDynamic(c *engine.Context) (interface{}, error) {
	return bilibot.GetDynamic(c.Query("mid"))
}
