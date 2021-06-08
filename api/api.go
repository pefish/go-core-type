package api

type IApi interface {
	GetDescription() string
	GetParamType() string
	GetParams() interface{}
}
