package env

import (
	"github.com/kelseyhightower/envconfig"
)

// APIConfig - APIサーバのホストとポートのコンフィグ
type APIConfig struct {
	Host string `split_words:"true"`
	Port string `split_words:"true"`
}

// ReadAPIEnv - APIサーバに関する設定を読み込む
func ReadAPIEnv() (APIConfig, error) {
	var APICfg APIConfig
	err := envconfig.Process("API", &APICfg)
	if err != nil {
		return APIConfig{}, err
	}
	return APICfg, nil
}
