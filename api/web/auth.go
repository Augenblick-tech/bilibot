package web

import (
	"time"

	"github.com/Augenblick-tech/bilibot/lib/engine"
	"github.com/Augenblick-tech/bilibot/lib/jwt"
	"github.com/Augenblick-tech/bilibot/pkg/dao"
	"github.com/Augenblick-tech/bilibot/pkg/e"
	"github.com/Augenblick-tech/bilibot/pkg/model"
	"github.com/Augenblick-tech/bilibot/pkg/model/api"
	"github.com/Augenblick-tech/bilibot/pkg/services/user"
)

// Register godoc
// @Summary      站点用户注册
// @Description
// @Tags         web
// @Accept       json
// @Produce      json
// @Param        information   body     api.UserInfo	true  "用户信息"
// @Router       /web/register [post]
func Register(c *engine.Context) (interface{}, error) {
	user := api.UserInfo{}

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
// @Tags         web
// @Accept       json
// @Produce      json
// @Param        SESSDATA   body     api.UserInfo	true  "用户信息"
// @Router       /web/login [post]
func Login(c *engine.Context) (interface{}, error) {
	tempUser := api.UserInfo{}

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

	token, err := jwt.GenToken(u.ID, u.Name)
	if err != nil {
		return nil, err
	}

	reToken, err := jwt.GenReToken(u.ID, u.Name)
	if err != nil {
		return nil, err
	}

	return api.RegisteredToken{
		BasicJWToken: api.BasicJWToken{
			AccessToken:         token,
			AccessTokenExpireAt: time.Now().Add(jwt.TokenExpireDuration).Unix(),
		},
		RefreshToken:         reToken,
		RefreshTokenExpireAt: time.Now().Add(jwt.ReTokenExpireDuration).Unix(),
	}, nil
}

// RefreshToken godoc
// @Summary      刷新 AccessToken
// @Description
// @Tags         web
// @Produce      json
// @Param 		 Authorization 	header 	string	true	"Bearer 刷新令牌"
// @Router       /web/refreshToken [get]
func RefreshToken(c *engine.Context) (interface{}, error) {
	token, err := jwt.ParseToken(c.Context.Request.Header.Get("Authorization"))
	if err != nil {
		return nil, err
	}
	if token.ExpiresAt.Unix() < time.Now().Unix() {
		return nil, e.RespCode_TokenExpired
	}
	accessToken, err := jwt.GenToken(token.UserID, token.Username)
	if err != nil {
		return nil, err
	}
	return api.BasicJWToken{
		AccessToken:         accessToken,
		AccessTokenExpireAt: time.Now().Add(jwt.TokenExpireDuration).Unix(),
	}, nil
}
