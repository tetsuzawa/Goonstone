package webcfg

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

// FRONTENDConfig - FRONTENDサーバのホストとポートのコンフィグ
type FRONTENDConfig struct {
	Host string `split_words:"true"`
	Port string `split_words:"true"`
}

// ReadFRONTENDEnv - FRONTENDサーバに関する設定を読み込む
func ReadFRONTENDEnv(cfg *FRONTENDConfig) error {
	err := envconfig.Process("FRONTEND", cfg)
	if err != nil {
		return fmt.Errorf("envconfig.Process: %w", err)
	}
	if cfg.Host == "" {
		cfg.Host = "127.0.0.1"
	}
	if cfg.Port == "" {
		cfg.Port = "80"
	}
	return nil
}
