package api

import (
	"log"
	"net/http"
	"time"

	"github.com/Augenblick-tech/bilibot/pkg/model"
	"github.com/Augenblick-tech/bilibot/response"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	quit     = make(chan string)
	status   = false
	dynamics = make(map[string][]model.Dynamic)
	mids     = make(map[string]struct{})
)

func RefreshDynamic(c *gin.Context) {
	var r response.Response

	mid := c.Query("mid")

	if _, ok := mids[mid]; ok {
		r.Code = response.CodeParamError
		r.JSON(c, http.StatusBadRequest, "mid already exists", nil)
		return
	} else {
		mids[mid] = struct{}{}
		log.Println("add mid", mid)
	}

	ticker := time.NewTicker(time.Second * time.Duration(viper.GetInt("user.RefreshTime")))

	go func() {
		status = true
		for {
			select {
			case v := <-quit:
				if v == c.Query("mid") {
					log.Println("quit mid:", v)
					return
				}
			case <-ticker.C:
				log.Println(c.Query("mid"))
				temp, err := model.GetDynamic(c.Query("mid"))
				if err != nil {
					log.Println(err)
					// 处理错误
					return
				}
				dynamics[mid] = temp
				log.Println("refresh")
			}
		}
	}()

	r.JSON(c, http.StatusOK, "success", nil)
}

func GetLatestDynamic(c *gin.Context) {
	var r response.Response

	mid := c.Query("mid")

	if len(dynamics) == 0 {
		r.Code = response.CodeParamError
		r.JSON(c, http.StatusBadRequest, "no dynamic", nil)
		return
	}

	r.JSON(c, http.StatusOK, "success", dynamics[mid][0])
}

func GetStatus(c *gin.Context) {
	var r response.Response

	mid := c.Query("mid")

	if _, ok := mids[mid]; ok && status {
		r.JSON(c, http.StatusOK, "runing", nil)
		return
	} else {
		r.JSON(c, http.StatusOK, "stop", nil)
		return
	}
}

func StopRefreshDynamic(c *gin.Context) {
	var r response.Response

	mid := c.Query("mid")

	if status {
		quit <- mid
		delete(mids, mid)
		status = false
		log.Println("stop mid:", mid)
		r.JSON(c, http.StatusOK, "success", nil)
		return
	} else {
		r.Code = response.CodeParamError
		r.JSON(c, http.StatusBadRequest, "stop failed", nil)
		return
	}
}
