package api

type UserInfo struct {
	Name     string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type BasicJWToken struct {
	AccessToken         string `json:"token"`
	AccessTokenExpireAt int64  `json:"expire_at"`
}

type RegisteredToken struct {
	BasicJWToken
	RefreshToken         string `json:"refresh_token"`
	RefreshTokenExpireAt int64  `json:"refresh_expire_at"`
}

type AuthorInfo struct {
	Mid   string `json:"mid"`
	BotID string `json:"bot_id"`
}

type BiliAuthInfo struct {
	SESSDATA string `json:"SESSDATA" binding:"required"`
}

type BiliQrCodeInfo struct {
	OauthKey string `json:"oauthKey" binding:"required"`
}
