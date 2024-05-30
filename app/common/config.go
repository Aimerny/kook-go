package common

import (
	jsoniter "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
	"os"
)

type Config struct {
	BotToken string `json:"token"`
	Compress bool   `json:"compress"`
}

const configPath string = "conf.json"

var defaultConf = &Config{
	BotToken: "Your kook-go bot token",
	Compress: true,
}

func ReadConfig() *Config {
	configData, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Error("read config file failed. generating default file...")
			data, _ := jsoniter.MarshalIndent(defaultConf, "", "  ")
			confFile, _ := os.OpenFile(configPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
			confFile.Write(data)
			os.Exit(1)
		}
	}
	conf := &Config{}
	err = jsoniter.Unmarshal(configData, conf)
	if err != nil {
		log.Panicf("read config file failed, err: %e", err)
	}
	return conf
}
