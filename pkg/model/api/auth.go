package api

// Request
type UserInfo struct {
	Name     string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthorInfo struct {
	Mid   string `json:"mid"`
	BotID string `json:"bot_id"`
}

// Response
type BasicJWToken struct {
	AccessToken         string `json:"token"`
	AccessTokenExpireAt int64  `json:"expire_at"`
}

type RegisteredToken struct {
	BasicJWToken
	RefreshToken         string `json:"refresh_token"`
	RefreshTokenExpireAt int64  `json:"refresh_expire_at"`
}

type BiliQrCodeAuth struct {
	TS       int    `json:"ts"`
	Url      string `json:"url"`
	OauthKey string `json:"oauth_key"`
}

type BotInfo struct {
	BotID   uint   `json:"bot_id"`
	Name    string `json:"name"`
	IsLogin bool   `json:"is_login"`
	Face    string `json:"face"`
}
