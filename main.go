package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"lottery_weichat/configs"
	"lottery_weichat/router"
)

func Init() {
	configs.InitGlobalConfig()
}
func main() {
	config := configs.GetGlobalConfig()
	fmt.Println(config)

	Init()
	fmt.Println("start")

	r := router.SetRouter()
	fmt.Println("start")
	if err := r.Run(fmt.Sprintf(":%d", config.AppConfig.Port)); err != nil {
		logrus.Errorf("sever run err :%v", err)
	}
}
