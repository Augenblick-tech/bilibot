package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lonzzi/bilibot/response"
)

func Login(c *gin.Context) {
	var r response.Response

	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == "" || password == "" {
		r.Code = response.CodeParamError
		r.JSON(c, http.StatusBadRequest, "username or password is empty", nil)
		return
	}

	r.Code = response.CodeSuccess
	r.JSON(c, http.StatusOK, "login success", nil)
}
