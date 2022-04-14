package config

import (
	"fmt"
	"path/filepath"
	"proj1/pkg/logger"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	Listen struct {
		Host string `mapstructure:"host"`
		Port string `mapstructure:"port"`
	} `mapstructure:"listen"`
	PostgresDB struct {
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		Database string `mapstructure:"database"`
		Timeout  int    `mapstructure:"timeout"`
		MaxConns int    `mapstructure:"max_conns"`
	} `mapstructure:"pg_db"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		var configName string
		configName = "local"

		viper.SetConfigName(configName)
		viper.SetConfigType("yml")

		dirPath, err := filepath.Abs("./configs")
		if err != nil {
			logger.Fatal(fmt.Sprintf("fatal error config dir: %s \n", err))
		}
		viper.AddConfigPath(dirPath)
		err = viper.ReadInConfig()
		if err != nil {
			logger.Fatal(fmt.Sprintf("fatal error config file, dir path: %s, error: %s \n", err, dirPath))
		}

		instance = &Config{}

		err = viper.Unmarshal(instance)
		if err != nil {
			logger.Fatal(fmt.Sprintf("fatal parse config: %s \n", err))
		}
	})

	return instance
}
