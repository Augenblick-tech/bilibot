package api

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lonzzi/BiliUpDynamicBot/pkg/model"
	"github.com/lonzzi/BiliUpDynamicBot/response"
	"github.com/spf13/viper"
)

var (
	quit     = make(chan int)
	status   = false
	dynamics = make(map[int][]model.Dynamic)
	mids     = make(map[int]interface{})
)

func RefreshDynamic(c *gin.Context) {
	var r response.Response

	mid, err := strconv.Atoi(c.Query("mid"))
	if err != nil {
		r.Code = response.CodeParamError
		r.JSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if _, ok := mids[mid]; ok {
		r.Code = response.CodeParamError
		r.JSON(c, http.StatusBadRequest, "mid already exists", nil)
		return
	} else {
		mids[mid] = nil
		log.Println("add mid", mid)
	}

	ticker := time.NewTicker(time.Second * time.Duration(viper.GetInt("user.RefreshTime")))

	go func() {
		status = true
		for {
			select {
			case v := <-quit:
				if strconv.Itoa(v) == c.Query("mid") {
					log.Println("quit mid:", v)
					return
				}
			case <-ticker.C:
				log.Println(c.Query("mid"))
				temp, err := model.GetDynamic(c.Query("mid"))
				if err != nil {
					r.Code = response.CodeRefreshError
					r.JSON(c, http.StatusBadGateway, err.Error(), nil)
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

	mid, err := strconv.Atoi(c.Query("mid"))
	if err != nil {
		r.Code = response.CodeParamError
		r.JSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	r.JSON(c, http.StatusOK, "success", dynamics[mid][0])
}

func GetStatus(c *gin.Context) {
	var r response.Response

	mid, err := strconv.Atoi(c.Query("mid"))
	if err != nil {
		r.Code = response.CodeParamError
		r.JSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

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

	mid, err := strconv.Atoi(c.Query("mid"))
	if err != nil {
		r.Code = response.CodeParamError
		r.JSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

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
