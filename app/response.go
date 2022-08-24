package app

import (
    "bytes"
    jsoniter "github.com/json-iterator/go"
    "net/http"
)

type Response struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
    Data    any    `json:"data,omitempty"`
}

// processParam Process response data
func (r *Response) processParam(param any) {
    switch param.(type) {
    case error:
        r.Message = param.(error).Error()
        break
    case int:
        r.Code = param.(int)
    case nil:
        break
    default:
        r.Data = param
        break
    }
}

// SendResponse Send api response
func (c *BaseContext) SendResponse(params ...any) error {
    r := &Response{
        Message: "ok",
        Code:    http.StatusOK,
    }

    for _, param := range params {
        r.processParam(param)
    }

    buffer := &bytes.Buffer{}
    encoder := jsoniter.NewEncoder(buffer)
    encoder.SetEscapeHTML(false)
    _ = encoder.Encode(r)

    return c.JSONBlob(http.StatusOK, buffer.Bytes())
}
