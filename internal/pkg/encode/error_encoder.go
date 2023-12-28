package encode

import (
	"errors"
	"fmt"
	stdhttp "net/http"

	e "github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type ResponseError struct {
	Code    int    `json:"code"`
	Reason  string `json:"reason"`
	Message string `json:"message"`
}

func NewResponseError(code int, reason, message string) *ResponseError {
	return &ResponseError{
		Code:    code,
		Reason:  reason,
		Message: message,
	}
}

func (e *ResponseError) Error() string {
	return fmt.Sprintf("HTTPError code: %d message: %s reason: %s", e.Code, e.Message, e.Reason)
}

func FromError(err error) *ResponseError {
	if err == nil {
		return nil
	}

	// 创建一个 kratos 的 Error，然后判断传进来的 err 是否是这个类型
	// 是的话，就变成我们自定义的格式，返回一个 ResponseError
	if se := new(e.Error); errors.As(err, &se) {
		return NewResponseError(int(se.Code), se.Reason, se.Message)
	}

	// 返回默认的错误
	return &ResponseError{
		Code:    500,
		Reason:  "UNKNOWN",
		Message: "unknown request error",
	}
}

func ErrorEncoder(w stdhttp.ResponseWriter, r *stdhttp.Request, err error) {
	se := FromError(err)

	codec, _ := http.CodecForRequest(r, "Accept")
	body, err := codec.Marshal(se)

	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/"+codec.Name())
	w.WriteHeader(se.Code)
	w.Write(body)
}
