package encode

import (
	"github.com/go-kratos/kratos/v2/transport/http"
)

func ResponseEncoder(w http.ResponseWriter, r *http.Request, data interface{}) error {
	type Response struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}

	res := &Response{
		Code:    200,
		Data:    data,
		Message: "",
	}

	codec, _ := http.CodecForRequest(r, "Accept")
	msRes, err := codec.Marshal(res)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/"+codec.Name())
	w.Write(msRes)
	return nil
}
