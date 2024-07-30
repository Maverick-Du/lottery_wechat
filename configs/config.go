package configs

import (
	rlog "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"sync"
	"time"
)

var (
	globalConfig GlobalConfig
	once         sync.Once
)

type GlobalConfig struct {
	AppConfig AppConf `yaml:"app" mapstructure:"app"`
	LogConfig LogConf `yaml:"log" mapstructure:"log"`
	DbConfig  DbConf  `yaml:"db" mapstructure:"db"`
}

type AppConf struct {
	AppName string `yaml:"app_name" mapstructure:"app_name"`
	Version string `yaml:"version" mapstructure:"version"`
	Port    int    `yaml:"port" mapstructure:"port"`
	RunMod  string `yaml:"run_mod" mapstructure:"run_mod"`
}

type DbConf struct {
	Host        string `yaml:"host" mapstructure:"host"`
	Port        string `yaml:"port" mapstructure:"port"`
	Password    string `yaml:"password" mapstructure:"password"`
	DbName      string `yaml:"dbname" mapstructure:"dbname"`
	MaxIdleComm int    `yaml:"max_idle_comm" mapstructure:"max_idle_comm"`
	MaxOpenConn int    `yaml:"max_open_conn" mapstructure:"max_open_conn"`
	MaxIdleTime int    `yaml:"max_idle_time" mapstructure:"max_idle_time"`
	User        string `yaml:"user" mapstructure:"user"`
}

type LogConf struct {
	LogPattern string `yaml:"log_pattern" mapstructure:"log_pattern"`
	LogPath    string `yaml:"log_path" mapstructure:"log_path"`
	SaveDays   uint   `yaml:"save_days" mapstructure:"save"`
	Level      string `yaml:"level" mapstructure:"level"`
}

func GetGlobalConfig() *GlobalConfig {
	once.Do(readConf)
	return &globalConfig
}

func readConf() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")

	err := viper.ReadInConfig()
	if err != nil {
		panic("read config file err" + err.Error())
	}
	err = viper.Unmarshal(&globalConfig)
	if err != nil {
		panic("unmarshal config file err")
	}
}

func InitGlobalConfig() {
	config := GetGlobalConfig()
	level, err := logrus.ParseLevel(config.LogConfig.Level)
	if err != nil {
		panic("parse error")
	}
	logrus.SetFormatter(&logFormatter{
		logrus.TextFormatter{
			DisableColors:   true,
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
		},
	})
	logrus.SetReportCaller(true)
	logrus.SetLevel(level)
	switch globalConfig.LogConfig.LogPattern {
	case "stdout":
		logrus.SetOutput(os.Stdout)
	case "stderr":
		logrus.SetOutput(os.Stderr)
	case "file":
		logger, err := rlog.New(
			config.LogConfig.LogPath+"./%Y%m%d",
			rlog.WithRotationTime(time.Hour*24),
			rlog.WithRotationCount(config.LogConfig.SaveDays),
		)
		if err != nil {
			panic("log err")
		}
		logrus.SetOutput(logger)
	default:
		panic("log init err")
	}
}
