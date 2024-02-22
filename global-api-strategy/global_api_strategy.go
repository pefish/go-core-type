package global_api_strategy

import (
	_type "github.com/pefish/go-core-type/api-session"
	go_error "github.com/pefish/go-error"
)

type IGlobalApiStrategy interface {
	Init(param interface{})
	Execute(out _type.IApiSession, param interface{}) *go_error.ErrorInfo
	Name() string
	Description() string
	ErrorCode() uint64
	SetErrorCode(code uint64) IGlobalApiStrategy
	SetErrorMsg(msg string) IGlobalApiStrategy
	ErrorMsg() string
}
