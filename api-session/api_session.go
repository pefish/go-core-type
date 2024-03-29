package api_session

import (
	"github.com/pefish/go-core-type/api"
	go_logger "github.com/pefish/go-logger"
	"io"
	"net/http"
)

type StatusCode int

type IApiSession interface {
	SetPathVars(vars map[string]string)
	PathVars() map[string]string
	SetJwtBody(jwtBody map[string]interface{})
	JwtBody() map[string]interface{}
	SetUserId(userId uint64)
	UserId() uint64
	SetJwtHeaderName(headerName string)
	JwtHeaderName() string
	ScanParams(dest interface{}) error
	MustScanParams(dest interface{})
	AddDefer(defer_ func())
	Defers() []func()
	SetData(key string, data interface{})
	Data(key string) interface{}
	Redirect(url string)
	WriteJson(data interface{}) error
	SetHeader(key string, value string)
	WriteText(text string) error
	SetStatusCode(code StatusCode)
	Host() string
	Path() string
	Body() io.ReadCloser
	Method() string
	Header(name string) string
	RemoteAddress() string
	UrlParams() map[string]string
	FormValues() (map[string][]string, error)
	ReadJSON(jsonObject interface{}) error
	Api() api.IApi
	SetApi(api api.IApi)
	ResponseWriter() http.ResponseWriter
	SetResponseWriter(w http.ResponseWriter)
	Request() *http.Request
	SetRequest(r *http.Request)
	Params() map[string]interface{}
	SetParams(params map[string]interface{})
	OriginalParams() map[string]interface{}
	SetOriginalParams(originalParams map[string]interface{})
	SetLang(lang string)
	Lang() string
	SetClientType(clientType string)
	ClientType() string
	Logger() go_logger.InterfaceLogger
}
