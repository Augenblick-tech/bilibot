package bili

import (
	"net/http"

	"github.com/Augenblick-tech/bilibot/pkg/model"
	"github.com/Augenblick-tech/bilibot/response"
	"github.com/gin-gonic/gin"
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
