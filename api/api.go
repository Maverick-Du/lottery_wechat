package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"lottery_weichat/internal/service"
	"net/http"
)

func Hello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "hello")
}

func InitPrize(ctx *gin.Context) {
	req := service.InitPrizeList{}
	if err := ctx.BindJSON(&req); err != nil {
		logrus.Errorf("InitPrize err:%v", err)
		ctx.JSON(http.StatusBadRequest, 200)
		return
	}
	if err := service.AddPrize(req.ViewPrizeList); err != nil {
		logrus.Errorf("API|addprize:%v", err)
		ctx.JSON(http.StatusInternalServerError, 500)
		return
	}
	ctx.JSON(200, "success")
}
func GetPrizeInfo(ctx *gin.Context) {
	rsp := service.GetPrizeInfoRsp{}
	prizeList, err := service.GetPrizeList()
	if err != nil {
		logrus.Errorf("API|getprizeinfo:%v", err)
		ctx.JSON(http.StatusInternalServerError, 500)
		return
	}
	var count int = 0
	var total int64
	for _, prize := range prizeList {
		if prize.Total == 0 || (prize.Total > 0 && prize.Left > 0) {
			count++
			total += prize.Left
		}
	}
	rsp.PrizeTypeNum = count
	rsp.PrizeTotal = total
	ctx.JSON(200, rsp)
}
func Lottery(ctx *gin.Context) {
	res := service.GetWinner()

	ctx.JSON(200, res)
}
