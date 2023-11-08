package config

import (
	"log"

	"github.com/spf13/viper"
)

type Server struct {
	Port string `yaml:"Port"`
	Host string `yaml:"Host"`
	Mode string `yaml:"Mode"`
}

type Loger struct {
	Development bool   `yaml:"Development"`
	Level       string `yaml:"Level"`
	Encoding    string `yaml:"encoding"`
}

type Config struct {
	Server Server `yaml:"server"`
	Loger  Loger  `yaml:"loger"`
}

func Read(name string) *Config {
	viper.AddConfigPath("./config")
	viper.SetConfigName(name)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	var cfg *Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Error unmarshalling config file, %s", err)
	}

	return cfg
}
