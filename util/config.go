package util

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"
)

//Config 全局配置文件
var Cfg = NewConfig()

//配置信息结构体
type Config struct {
	Redis   RedisConfig `json:"redis"`
	Host    string      `json:"host"`
	Timeout time.Duration `json:"timeout"`
	Cron    CronConfig `json:"crons"`
}

type RedisConfig struct {
	Addr      string `json:"addr"`
	Key       string `json:"key"`
	MaxIdle   int    `json:"maxIdle"`
	MaxActive int    `json:"maxActive"`
}

//定时任务
type CronConfig struct {
	GetProxyCron   string `json:"getProxyCron"`
	CheckRedisCron string `json:"checkRedisCron"`
}

//创建配置文件
func NewConfig() *Config {
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatalln("Read config file error.")
	}
	config := &Config{}
	err = json.Unmarshal(file, config)
	if err != nil {
		panic(err)
	}

	return config
}
