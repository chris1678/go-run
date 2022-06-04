package response

import (
	"github.com/chris1678/go-run/utils"
	"net/http"
	"time"
)

// APIException api结构体
type APIException struct {
	Code      int         `json:"code"`
	Success   bool        `json:"success"`
	Msg       string      `json:"msg"`
	Timestamp string      `json:"timestamp"`
	Result    interface{} `json:"result"`
}

// 实现接口
func (e *APIException) Error() string {
	return e.Msg
}

func newAPIException(code int, msg string, data interface{}, success bool) *APIException {
	return &APIException{
		Code:      code,
		Success:   success,
		Msg:       msg,
		Timestamp: time.Now().Format(utils.TimeFormat),
		Result:    data,
	}
}

// ServerError 500 错误处理
func ServerError() *APIException {
	return newAPIException(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), nil, false)
}

// NotFound 404 错误
func NotFound() *APIException {
	return newAPIException(http.StatusNotFound, http.StatusText(http.StatusNotFound), nil, false)
}

// UnknownError 未知错误
func UnknownError(message string) *APIException {
	return newAPIException(http.StatusForbidden, message, nil, false)
}

// ParameterError 参数错误
func ParameterError(message string) *APIException {
	return newAPIException(http.StatusBadRequest, message, nil, false)
}

// AuthError 授权错误
func AuthError(message string) *APIException {
	return newAPIException(http.StatusBadRequest, message, nil, false)
}

// ResponseJson 200
func ResponseJson(message string, data interface{}, success bool) *APIException {
	return newAPIException(http.StatusOK, message, data, success)
}
