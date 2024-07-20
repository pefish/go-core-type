package api

type IApi interface {
	Description() string
	ParamType() string
	Params() interface{}
}

type ApiResult struct {
	Msg  string      `json:"msg"`
	Code uint64      `json:"code"`
	Data interface{} `json:"data"`
}
