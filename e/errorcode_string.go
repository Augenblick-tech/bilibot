// Code generated by "stringer -type=ErrorCode --linecomment"; DO NOT EDIT.

package e

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ERR_READFILE-0]
	_ = x[ERR_WRITEFILE-1]
	_ = x[ERR_CREATEFILE-2]
	_ = x[ERR_UNMARSHAL-3]
	_ = x[ERR_MARSHAL-4]
	_ = x[ERR_ADD_DYNAMIC-5]
	_ = x[ERR_REPLY_DYNAMIC-6]
	_ = x[ERR_DYNAMIC_EXIST-7]
	_ = x[ERR_INVALID_NUMBER-8]
	_ = x[ERR_BELOW_THRESHOLD-9]
	_ = x[ERR_COMMENT_REPLY_FAIL-10]
	_ = x[ERR_HTTP_STATUS_NOT_OK-11]
	_ = x[ERR_LOGIN_FAIL-12]
}

const _ErrorCode_name = "读取文件错误写入文件错误创建文件错误反序列化错误序列化错误添加动态错误回复动态错误动态已存在无效数字低于阈值评论错误http状态不是200登录失败"

var _ErrorCode_index = [...]uint8{0, 18, 36, 54, 72, 87, 105, 123, 138, 150, 162, 174, 193, 205}

func (i ErrorCode) String() string {
	if i < 0 || i >= ErrorCode(len(_ErrorCode_index)-1) {
		return "ErrorCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ErrorCode_name[_ErrorCode_index[i]:_ErrorCode_index[i+1]]
}