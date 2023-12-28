package encode

import (
	"errors"
	"fmt"
	stdhttp "net/http"

	e "github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type ResponseError struct {
	Code   int    `json:"code"`
	Reason string `json:"reason"`
	Msg    string `json:"msg"`
}

func NewResponseError(code int, reason, message string) *ResponseError {
	return &ResponseError{
		Code:   code,
		Reason: reason,
		Msg:    message,
	}
}

func (e *ResponseError) Error() string {
	return fmt.Sprintf("ResponseError code: %d message: %s reason: %s", e.Code, e.Msg, e.Reason)
}

func FromError(err error) *ResponseError {
	if err == nil {
		return nil
	}

	if se := new(e.Error); errors.As(err, &se) {
		return NewResponseError(int(se.Code), se.Reason, se.Message)
	}

	return &ResponseError{
		Code:   500,
		Reason: "UNKNOWN",
		Msg:    "unknown request error",
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
