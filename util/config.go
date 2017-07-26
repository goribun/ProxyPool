package util

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Config struct defines the config structure
type Config struct {
	Redis RedisConfig `json:"redis"`
	Host  string      `json:"host"`
}

type RedisConfig struct {
	Addr      string `json:"addr"`
	Key       string  `json:"key"`
	MaxIdle   int `json:"maxIdle"`
	MaxActive int `json:"maxActive"`
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
