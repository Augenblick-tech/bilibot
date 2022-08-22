package api

// Requtest
type CommentInfo struct {
	BotID   string `json:"bot_id" binding:"required"`
	Type    int    `json:"type" binding:"required"`
	Oid     string `json:"oid" binding:"required"`
	Message string `json:"message" binding:"required"`
}

// Response
type BasicResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type ReplyInfo struct {
	SuccessToast string      `json:"success_toast"`
	Emote        interface{} `json:"emote"`
}
