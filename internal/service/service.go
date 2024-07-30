package service

import (
	"github.com/sirupsen/logrus"
	"lottery_weichat/constant"
	"lottery_weichat/internal/model"
	"lottery_weichat/internal/repo"
	"math/rand"
	"time"
)

func AddPrize(viewPrizeList []*ViewPrize) error {
	prizeList := make([]*model.Prize, 0)
	for _, viewPrize := range viewPrizeList {
		prize := &model.Prize{
			ID:             viewPrize.ID,
			Name:           viewPrize.Name,
			Pic:            viewPrize.Pic,
			Left:           viewPrize.Left,
			Link:           viewPrize.Link,
			Type:           viewPrize.Type,
			Data:           viewPrize.Data,
			Total:          viewPrize.Total,
			IsUse:          viewPrize.IsUse,
			Probability:    viewPrize.Probability,
			ProbabilityMax: viewPrize.ProbabilityMax,
			ProbabilityMin: viewPrize.ProbabilityMin,
		}
		prizeList = append(prizeList, prize)
	}
	if err := repo.AddPrize(prizeList); err != nil {
		logrus.Errorf("service|Addpize :%v", err)
		return err
	}

	return nil
}

func GetPrizeList() ([]*model.Prize, error) {
	prizeList, err := repo.GetPrizeList()
	if err != nil {
		logrus.Errorf("service|Getpize :%v", err)
	}
	return prizeList, err
}

func GetWinner() map[string]interface{} {
	code := LunkyCode()
	logrus.Infof("service|GetWinner :%d\n", code)
	var ok bool
	res := make(map[string]interface{})
	res["中奖信息"] = "未中奖"

	prizeList, err := repo.GetPrizeList()
	if err != nil {
		logrus.Errorf("service|Getpize :%v", err)
		return nil
	}
	for _, prize := range prizeList {
		if prize.IsUse != constant.PrizeInUse || (prize.Total > 0 && prize.Left <= 0) {
			continue
		}
		if prize.ProbabilityMin <= int64(code) && prize.ProbabilityMax > int64(code) {
			var profile string
			ok, profile = PrizeSenderMap[int(prize.Type)].SendPrize(prize)
			if ok {
				delete(res, "中奖信息")
				res["success"] = ok
				res["id"] = prize.ID
				res["name"] = prize.Name
				res["profile"] = profile
				res["link"] = prize.Link
				break

			}
		}
	}
	return res
}
func LunkyCode() int32 {
	seed := time.Now().UnixNano()
	code := rand.New(rand.NewSource(seed)).Int31n(constant.ProbalityLimit) //返回0-9999的随机数
	return code
}
