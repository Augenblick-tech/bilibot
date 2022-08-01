package api

import (
	"log"
	"time"

	"github.com/Augenblick-tech/bilibot/lib/engine"
	"github.com/Augenblick-tech/bilibot/pkg/e"
	"github.com/Augenblick-tech/bilibot/pkg/model"
	"github.com/spf13/viper"
)

var (
	quit     = make(chan string)
	status   = false
	dynamics = make(map[string][]model.Dynamic)
	mids     = make(map[string]struct{})
)

func RefreshDynamic(c *engine.Context) (interface{}, error) {

	mid := c.Context.Query("mid")

	if _, ok := mids[mid]; ok {
		return nil, e.RespCode_RefreshError
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
				if v == mid {
					log.Println("quit mid:", v)
					return
				}
			case <-ticker.C:
				log.Println(mid)
				temp, err := model.GetDynamic(mid)
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

	return "success", nil
}

func GetLatestDynamic(c *engine.Context) (interface{}, error) {
	mid := c.Context.Query("mid")

	if len(dynamics) == 0 {
		return nil, e.RespCode_ParamError
	}

	return dynamics[mid][0], nil
}

func GetStatus(c *engine.Context) (interface{}, error) {
	mid := c.Context.Query("mid")

	if _, ok := mids[mid]; ok && status {
		return "running", nil
	} else {
		return "stop", nil
	}
}

func StopRefreshDynamic(c *engine.Context) (interface{}, error) {
	mid := c.Context.Query("mid")

	if status {
		quit <- mid
		delete(mids, mid)
		status = false
		log.Println("stop mid:", mid)
		return "success", nil
	} else {
		return "stop failed", nil
	}
}
