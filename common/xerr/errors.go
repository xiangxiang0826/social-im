package xerr

import (
	"fmt"
)

/**
常用通用固定错误
*/

type CodeError struct {
	errCode uint32
	errMsg  string
}

//返回给前端的错误码
func (e *CodeError) GetErrCode() uint32 {
	return e.errCode
}

//返回给前端显示端错误信息
func (e *CodeError) GetErrMsg() string {
	return e.errMsg
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("ErrCode:%d，ErrMsg:%s", e.errCode, e.errMsg)
}

func NewErrCodeMsg(errCode uint32, errMsg string) *CodeError {
	return &CodeError{errCode: errCode, errMsg: errMsg}
}

func NewErrCode(errCode uint32) *CodeError {
	errMsg := MapErrMsg(errCode)
	if errMsg == "" {
		errMsg = GetCommomErrorMsg()
	}
	return &CodeError{errCode: errCode, errMsg: errMsg}
}

func NewErrMsg(errMsg string) *CodeError {
	return &CodeError{errCode: SERVER_COMMON_ERROR, errMsg: errMsg}
}

func NewErrWithFormatMsg(errCode uint32, format string, args ...interface{}) *CodeError {
	preMsg := fmt.Sprintf(format, args...)
	errMsg := preMsg +" "+ MapErrMsg(errCode)
	return &CodeError{errCode: errCode, errMsg: errMsg}
}

func NewErrWithCodeMsg(errCode uint32, errMsg string) *CodeError {
	return &CodeError{errCode: errCode, errMsg: errMsg}
}
