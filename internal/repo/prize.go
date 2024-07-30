package repo

import (
	"github.com/sirupsen/logrus"
	"lottery_weichat/constant"
	"lottery_weichat/internal/model"
	"lottery_weichat/internal/pkg/gormcli"
)

func AddPrize(prizeList []*model.Prize) error {
	db := gormcli.GetDb()
	if err := db.Model(&model.Prize{}).Create(prizeList).Error; err != nil {
		logrus.Errorf("repo|add err:%v", err)
		return err
	}
	logrus.Infof("repo|add prize success")
	return nil
}
func GetPrizeList() ([]*model.Prize, error) {
	db := gormcli.GetDb()
	var prizeList []*model.Prize
	if err := db.Model(&model.Prize{}).Where("is_use= ?", constant.PrizeInUse).Find(&prizeList).Error; err != nil {
		logrus.Errorf("repo|get prize list err:%v", err)
		return nil, err
	}
	return prizeList, nil
}
func SavePrize(prize *model.Prize) error {
	db := gormcli.GetDb()
	if err := db.Model(&model.Prize{}).Where("id= ?", prize.ID).Save(prize).Error; err != nil {
		logrus.Errorf("repo|save prize err:%v", err)
		return err
	}
	logrus.Infof("repo|save prize success")
	return nil
}
