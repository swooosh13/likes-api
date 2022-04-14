package config

import (
	"fmt"
	"proj1/pkg/logger"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	IsDebug *bool `yaml:"is_debug"`
	Listen  struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"app"`
	PostgresDB struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		DBname string `yaml:"dbname"`
		Timeout int `yaml:"timeout"`
	} `yaml:"pgdb"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		viper.SetConfigName("local")
		viper.SetConfigType("yml")
		viper.AddConfigPath("./configs")
		err := viper.ReadInConfig()
		if err != nil {
			logger.Fatal(fmt.Sprint("fatal error config file: %w \n", err))
		}

		instance = &Config{}

		err = viper.Unmarshal(instance)
		if err != nil {
			logger.Fatal(fmt.Sprint("Fatal parse config: %w \n", err))
		}
	})

	return instance
}
