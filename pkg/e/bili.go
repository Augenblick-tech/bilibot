package e

// bili scan code
const (
	NotConfirmed     ErrCode = -5 + iota // 已扫码，未确认
	Waiting                              // 等待扫码
	ErrBiliUndefined                     // 未定义
	KeyTimeout                           // 秘钥超时
	KeyInvalid                           // 秘钥无效
)

// bili response code
const (
// ...
)
