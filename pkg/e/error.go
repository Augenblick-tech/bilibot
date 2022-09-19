package e

type ErrCode int

//go:generate stringer -type=ErrCode --linecomment
const (
	ErrUnmarshal       ErrCode = 50000 + iota // 反序列化错误
	ErrMarshal                                // 序列化错误
	ErrAddDynamic                             // 添加动态错误
	ErrReplyFailed                            // 回复错误
	ErrExisted                                // 数据已存在
	ErrInvalidNumber                          // 无效数字
	ErrLoginFailed                            // 登录失败
	ErrInvalidPassword                        // 密码错误
	ErrEmptyAuth                              // 请求体中auth为空
	ErrFormat                                 // 格式错误
	ErrInvalidToken                           // Token非法
	ErrInvalidParam                           // 参数非法
	ErrTokenExpired                           // token已过期
	ErrNotFound                               // 没有找到此条记录
	ErrCreate                                 // 创建失败
	ErrBinding                                // 绑定失败
	ErrNotLogin                               // 未登录
)

func (e ErrCode) Error() string {
	return e.String()
}
