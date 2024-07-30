package router

import (
	"github.com/gin-gonic/gin"
	"lottery_weichat/api"
)

func SetRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	group := r.Group("/lottery_wechat")
	group.GET("/hello", api.Hello)
	group.POST("/add_prize", api.InitPrize)
	group.GET("/get_prize_info", api.GetPrizeInfo)
	group.GET("/lottery", api.Lottery)

	return r
}
