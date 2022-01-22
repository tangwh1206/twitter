package faults

var (
	CodeSuccess          Code = 0
	CodeUserNotFound     Code = 1001
	CodePasswordError    Code = 1002
	CodeOperatorDeny     Code = 1003
	CodeUserAlreadyExist Code = 1004

	CodeBadRequest    Code = 2001
	CodeInternelError Code = 2002
	CodeSystemBusy    Code = 2003
)

var (
	MsgSuccess          = "成功"
	MsgUserNotFound     = "用户未注册"
	MsgBadRequest       = "请求格式错误"
	MsgOperateDeny      = "用户无操作权限"
	MsgInternalError    = "系统内部错误"
	MsgPasswordError    = "密码输入错误"
	MsgUserAlreadyExist = "用户名称已存在"
	MsgSystemBusy       = "系统繁忙，请稍后重试"
)

var (
	ErrUserNotFound     = New(CodeUserNotFound, MsgUserNotFound, nil)
	ErrBadRequest       = New(CodeBadRequest, MsgBadRequest, nil)
	ErrOperateDeny      = New(CodeOperatorDeny, MsgOperateDeny, nil)
	ErrInternalError    = New(CodeInternelError, MsgInternalError, nil)
	ErrPasswordError    = New(CodePasswordError, MsgPasswordError, nil)
	ErrUserAlreadyExist = New(CodeUserAlreadyExist, MsgUserAlreadyExist, nil)
	ErrSystemBusy       = New(CodeSystemBusy, MsgSystemBusy, nil)
)
