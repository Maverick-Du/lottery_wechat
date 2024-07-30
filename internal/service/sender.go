package service

import (
	"github.com/sirupsen/logrus"
	"lottery_weichat/constant"
	"lottery_weichat/internal/model"
	"lottery_weichat/internal/repo"
)

type PrizeSender interface {
	SendPrize(prize *model.Prize) (bool, string)
}
type CoinSender struct {
}

func (c *CoinSender) SendPrize(prize *model.Prize) (bool, string) {
	if prize.Total == 0 {
		return true, prize.Data
	}
	if prize.Left <= 0 {
		return false, "没有奖品了"
	}
	prize.Left--
	if err := repo.SavePrize(prize); err != nil {
		logrus.Errorf("save prize error:%v", err)
		return false, err.Error()
	}
	return true, prize.Data
}

type CouponSender struct {
}

func (c *CouponSender) SendPrize(prize *model.Prize) (bool, string) {
	if prize.Total == 0 {
		return true, prize.Data
	}
	if prize.Left <= 0 {
		return false, "没有奖品了"
	}
	prize.Left--
	if err := repo.SavePrize(prize); err != nil {
		logrus.Errorf("save prize error:%v", err)
		return false, err.Error()
	}
	return true, prize.Data
}

type SmallEntitySender struct {
}

func (c *SmallEntitySender) SendPrize(prize *model.Prize) (bool, string) {
	if prize.Total == 0 {
		return true, prize.Data
	}
	if prize.Left <= 0 {
		return false, "没有奖品了"
	}
	prize.Left--
	if err := repo.SavePrize(prize); err != nil {
		logrus.Errorf("save prize error:%v", err)
		return false, err.Error()
	}
	return true, prize.Data
}

type BigEntitySender struct {
}

func (c *BigEntitySender) SendPrize(prize *model.Prize) (bool, string) {
	if prize.Total == 0 {
		return true, prize.Data
	}
	if prize.Left <= 0 {
		return false, "没有奖品了"
	}
	prize.Left--
	if err := repo.SavePrize(prize); err != nil {
		logrus.Errorf("save prize error:%v", err)
		return false, err.Error()
	}
	return true, prize.Data
}

var PrizeSenderMap = map[int]PrizeSender{
	constant.PrizeTypeCoin:        &CoinSender{},
	constant.PirzeTypeCoupon:      &CouponSender{},
	constant.PrizeTypeSmallEntity: &SmallEntitySender{},
	constant.PrizeTypeBigEntity:   &BigEntitySender{},
}
