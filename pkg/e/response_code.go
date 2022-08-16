package e

type ResponseCode int

//go:generate stringer -type=ResponseCode --linecomment
const (
	RespCode_Success         ResponseCode = iota // 成功
	RespCode_ParamError                          // 参数错误
	RespCode_LoginTimeout                        // 登录超时
	RespCode_LoginError                          // 登录错误
	RespCode_GetDynamicError                     // 获取动态错误
	RespCode_ReplyError                          // 回复错误
	RespCode_RefreshError                        // 刷新错误
	RespCode_AlreadyExist                        // 已存在
	RespCode_NO_RECORD                           // 当前用户无此up主/bot数据
)

func (r ResponseCode) Error() string {
	return r.String()
}
