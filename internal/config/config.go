package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	IsDebug *bool `yaml:"is_debug"`
	App     struct {
		BindIP string `yaml:"bind_ip"`
		Port   string `yaml:"port"`
	} `yaml:"app"`
	PgDB struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"pgdb"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		log.Println("Read config")
		instance = &Config{}

		if err := cleanenv.ReadConfig("configs/local.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			log.Println(help)
			log.Fatal(err)
		}
	})

	return instance
}
