package engine

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Result(h Handle) Handle {
	return func(c *Context) (result interface{}, err error) {
		if r, err := h(c); err != nil {
			JsonError(c, r, err)
		} else {
			JsonResult(c, r)
		}
		return nil, nil
	}
}

type Error struct {
	Code string
	Err  error
}

func (e Error) Error() string {
	return e.Err.Error()
}

func JsonError(ctx *Context, data interface{}, err error) {
	code := "500"
	e := err
	if es, ok := err.(Error); ok {
		code = es.Code
		e = es.Err
	}
	ctx.Context.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
		"msg":  e.Error(),
	})
	ctx.Context.Abort()
}

func JsonResult(ctx *Context, data interface{}) {
	if data == nil {
		data = "ok"
	}
	ctx.Context.JSON(http.StatusOK, gin.H{
		"code": "200",
		"data": data,
	})
}