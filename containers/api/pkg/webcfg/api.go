package webcfg

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

// APIConfig - APIサーバのホストとポートのコンフィグ
type APIConfig struct {
	Host string `split_words:"true"`
	Port string `split_words:"true"`
}

// ReadAPIEnv - APIサーバに関する設定を読み込む
func ReadAPIEnv(prefix string, cfg *APIConfig) error {
	var err error
	if prefix == "" {
		err = envconfig.Process("API", cfg)
	} else {
		err = envconfig.Process(prefix+"_API", cfg)
	}
	if err != nil {
		return fmt.Errorf("envconfig.Process: %w", err)
	}
	if cfg.Host == "" {
		cfg.Host = "127.0.0.1"
	}
	if cfg.Port == "" {
		cfg.Port = "8080"
	}
	return nil
}
