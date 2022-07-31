package bili

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lonzzi/bilibot/pkg/model"
	"github.com/lonzzi/bilibot/response"
)

func AddReply(c *gin.Context) {
	var r response.Response

	replyResp, err := model.AddReply(c.PostForm("type"), c.PostForm("oid"), c.PostForm("message"))
	if err != nil {
		r.Code = response.CodeReplyError
		r.JSON(c, http.StatusBadGateway, err.Error(), nil)
		return
	}

	r.JSON(c, http.StatusOK, "success", replyResp)
}
