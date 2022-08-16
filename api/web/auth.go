package web

import (
	"github.com/Augenblick-tech/bilibot/lib/engine"
	"github.com/Augenblick-tech/bilibot/lib/jwt"
	"github.com/Augenblick-tech/bilibot/pkg/dao"
	"github.com/Augenblick-tech/bilibot/pkg/e"
	"github.com/Augenblick-tech/bilibot/pkg/model"
	"github.com/Augenblick-tech/bilibot/pkg/services/user"
)

type userInfo struct {
	Name     string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Register godoc
// @Summary      站点用户注册
// @Description
// @Tags         v2
// @Accept       json
// @Produce      json
// @Param        information   body     userInfo	true  "用户信息"
// @Router       /web/register [post]
func Register(c *engine.Context) (interface{}, error) {
	user := userInfo{}

	err := c.Bind(&user)
	if err != nil {
		return nil, err
	}

	// password encryption

	err = dao.Create(&model.User{
		Name:     user.Name,
		Password: user.Password,
	})

	return user.Name, err
}

// Login godoc
// @Summary      站点用户登录
// @Description
// @Tags         v2
// @Accept       json
// @Produce      json
// @Param        SESSDATA   body     userInfo	true  "用户信息"
// @Router       /web/login [post]
func Login(c *engine.Context) (interface{}, error) {
	tempUser := userInfo{}

	err := c.Bind(&tempUser)
	if err != nil {
		return nil, err
	}

	u, err := user.Get(tempUser.Name)
	if err != nil {
		return nil, err
	}

	// password decryption

	if u.Password != tempUser.Password {
		return nil, e.RespCode_ParamError
	}

	return jwt.GenToken(u.ID, u.Name)
}
