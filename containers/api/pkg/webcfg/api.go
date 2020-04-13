package webcfg

import "github.com/kelseyhightower/envconfig"

// APIConfig - APIサーバのホストとポートのコンフィグ
type APIConfig struct {
	Host string `split_words:"true"`
	Port string `split_words:"true"`
}

// ReadAPIEnv - APIサーバに関する設定を読み込む
func ReadAPIEnv(cfg *APIConfig) (*APIConfig, error) {
	err := envconfig.Process("API", cfg)
	if err != nil {
		return cfg, err
	}
	if cfg.Host == "" {
		cfg.Host = "127.0.0.1"
	}
	if cfg.Port == "" {
		cfg.Port = "8080"
	}
	return cfg, nil
}
