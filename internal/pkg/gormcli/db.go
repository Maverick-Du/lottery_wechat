package gormcli

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"lottery_weichat/configs"
	"sync"
	"time"
)

var (
	db   *gorm.DB
	once sync.Once
)

func openDb() {
	dbConfig := configs.GetGlobalConfig().DbConfig

	conbArgs := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DbName)
	logrus.Infof("connArgs:%s", conbArgs)

	var err error
	db, err = gorm.Open(mysql.Open(conbArgs), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("fail to conn db:%v", err))
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic("fetch db err:" + err.Error())
	}
	sqlDB.SetMaxIdleConns(dbConfig.MaxIdleTime)
	sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConn)
	sqlDB.SetConnMaxIdleTime(time.Duration(dbConfig.MaxIdleTime * int(time.Second)))

}

func GetDb() *gorm.DB {
	once.Do(openDb)
	return db
}
