package api

type IApi interface {
	Description() string
	ParamType() string
	Params() interface{}
}
