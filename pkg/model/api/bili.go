package api

type CommentInfo struct {
	BotID   string `json:"bot_id" binding:"required"`
	Type    int    `json:"type" binding:"required"`
	Oid     string `json:"oid" binding:"required"`
	Message string `json:"message" binding:"required"`
}
