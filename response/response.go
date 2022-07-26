package response

import "github.com/gin-gonic/gin"

type Response struct {
	Code int
}

func (r *Response) JSON(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(code, gin.H{
		"code": r.Code,
		"msg":  msg,
		"data": data,
	})
}
