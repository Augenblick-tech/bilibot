package bili

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lonzzi/BiliUpDynamicBot/pkg/model"
	"github.com/lonzzi/BiliUpDynamicBot/response"
)

func GetDynamic(c *gin.Context) {
	var r response.Response

	dynamics, err := model.GetDynamic(c.Query("mid"))
	if err != nil {
		r.Code = response.CodeGetDynamicError
		r.JSON(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	r.JSON(c, http.StatusOK, "success", dynamics)
}