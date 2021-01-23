package Response

import "github.com/gin-gonic/gin"

type Response struct {
	Code int          `json:"code"`
	Data interface{}  `json:"data"`
	Msg  interface{}  `json:"msg"`
	C    *gin.Context `json:"-"`
}

func NewResponse(c *gin.Context) *Response {
	return &Response{C: c}
}

func (r *Response) Show(httpCode int, data interface{}, msg interface{}) {
	r.C.JSON(httpCode, Response{Code: httpCode, Data: data, Msg: msg})
	return
}
