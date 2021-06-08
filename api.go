package go_core_type

type IApi interface {
	GetDescription() string
	GetParamType() string
	GetParams() interface{}
}
