package redisx

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
	"github.com/kelseyhightower/envconfig"
)

// Config - Redisの接続情報に関する設定
type Config struct {
	Protocol string `split_words:"true"`
	Host     string `split_words:"true"`
	Port     string `split_words:"true"`
}

// ReadEnv - 指定したenvfileからRedisに関する設定を読み込む
func ReadEnv(cfg *Config) (*Config, error) {
	err := envconfig.Process("REDIS", cfg)
	if err != nil {
		return cfg, err
	}
	return cfg, nil
}

func (c Config) build() Config {
	if c.Protocol == "" {
		c.Protocol = "tcp"
	}
	if c.Host == "" {
		c.Host = "127.0.0.1"
	}
	if c.Port == "" {
		c.Port = "6379"
	}
	return c
}

// Connect - Redisに接続
func Connect(c Config) (redis.Conn, error) {
	c = c.build()

	conn, err := redis.Dial(c.Protocol, fmt.Sprintf("%s:%s", c.Host, c.Port))
	if err != nil {
		return nil, err
	}
	return conn, nil
}
