package e

type LoginCode int

//go:generate stringer -type=LoginCode --linecomment
const (
	NOT_CONFIRM LoginCode = -5 + iota // 已扫码，未确认
	NOT_SCAN                          //等待扫码
	KEY_TIMEOUT LoginCode = -2        // 秘钥超时
	KEY_INVALID LoginCode = -1        // 秘钥无效
)

func (l LoginCode) Error() string {
	return l.String()
}