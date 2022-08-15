package e

type ErrorCode int

//go:generate stringer -type=ErrorCode --linecomment
const (
	ERR_READFILE           ErrorCode = iota // 读取文件错误
	ERR_WRITEFILE                           // 写入文件错误
	ERR_CREATEFILE                          // 创建文件错误
	ERR_UNMARSHAL                           // 反序列化错误
	ERR_MARSHAL                             // 序列化错误
	ERR_ADD_DYNAMIC                         // 添加动态错误
	ERR_REPLY_DYNAMIC                       // 回复动态错误
	ERR_DYNAMIC_EXIST                       // 动态已存在
	ERR_INVALID_NUMBER                      // 无效数字
	ERR_BELOW_THRESHOLD                     // 低于阈值
	ERR_COMMENT_REPLY_FAIL                  // 评论错误
	ERR_HTTP_STATUS_NOT_OK                  // http状态不是200
	ERR_LOGIN_FAIL                          // 登录失败
	ERR_AUTH_EMPTY                          // 请求体中auth为空
	ERR_AUTH_FORMAT                         // 请求体中auth格式错误
)

func (e ErrorCode) Error() string {
	return e.String()
}
