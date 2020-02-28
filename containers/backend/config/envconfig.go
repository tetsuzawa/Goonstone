package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type APIConfig struct {
	Host string `split_words:"true"`
	Port string `split_words:"true"`
}

var API APIConfig

func init()  {
	err := envconfig.Process("", &API)
	if err != nil {
		log.Fatal(err.Error())
	}
}
