package message

import (
	"git.tvblack.com/video/frame/proto/p_common"
)

type Result struct {
	Code    p_common.RequestCode
	ErrCode int32
	Msg     string
}

func (r *Result) Success() bool {
	return r.Code == p_common.RequestCode_RequestCodeSuccess
}

func (r *Result) IsError() bool {
	return r.Code != p_common.RequestCode_RequestCodeSuccess
}

func (r *Result) SetError(code p_common.RequestCode, msg string) {
	r.Code = code
	r.Msg = msg
}

func NewResult() *Result {
	return &Result{
		Code:    p_common.RequestCode_RequestCodeSuccess,
		ErrCode: 0,
		Msg:     "",
	}
}

func Error(code p_common.RequestCode, eCode int32, msg string) *Result {
	return &Result{
		Code:    code,
		ErrCode: eCode,
		Msg:     msg,
	}
}
