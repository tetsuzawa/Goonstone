package config

import (
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
)

type APIConfig struct {
	Host string `split_words:"true"`
	Port string `split_words:"true"`
}

type DBConfig struct {
	User       string `split_words:"true"`
	Password   string `split_words:"true"`
	Host       string `split_words:"true"`
	Port       string `split_words:"true"`
	Database   string `split_words:"true"`
	GormPrefix string `split_words:"true"`
}

var API APIConfig
var DB DBConfig

func init() {
	err := envconfig.Process("API", &API)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = envconfig.Process("DB", &API)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(API.Port)
	fmt.Println(DB.Password)
}
