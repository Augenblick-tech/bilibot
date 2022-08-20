package e

// bili scan code
const (
	NotConfirmed ErrCode = -5 + iota // 已扫码，未确认
	Waiting                           // 等待扫码
	ErrUnkonwn                        // 未知错误
	KeyTimeout                        // 秘钥超时
	KeyInvalid                        // 秘钥无效
)

// bili response code
const (
	// ...
)
