package api_strategy

import (
	_type "github.com/pefish/go-core-type/api-session"
	go_error "github.com/pefish/go-error"
)

type IApiStrategy interface {
	Execute(out _type.IApiSession, param interface{}) *go_error.ErrorInfo
	GetName() string
	GetDescription() string
	GetErrorCode() uint64
	SetErrorCode(code uint64) IApiStrategy
	SetErrorMsg(msg string) IApiStrategy
	GetErrorMsg() string
}