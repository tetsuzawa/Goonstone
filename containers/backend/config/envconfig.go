package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type APIConfig struct {
	Host string `split_words:"true"`
	Port string `split_words:"true"`
}

type DBConfig struct {
	User       string `required:"true" split_words:"true"`
	Password   string `required:"true" split_words:"true"`
	Host       string `required:"true" split_words:"true"`
	Port       string `required:"true" split_words:"true"`
	Database   string `required:"true" split_words:"true"`
	GormPrefix string `required:"true" split_words:"true"`
}

var API APIConfig
var DB DBConfig

func init() {
	err := envconfig.Process("API", &API)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = envconfig.Process("DB", &DB)
	if err != nil {
		log.Fatal(err.Error())
	}
}
