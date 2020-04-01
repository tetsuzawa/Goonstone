package env

import (
	"github.com/kelseyhightower/envconfig"
)

// FRONTENDConfig - FRONTENDサーバのホストとポートのコンフィグ
type FRONTENDConfig struct {
	Host string `split_words:"true"`
	Port string `split_words:"true"`
}

// ReadFRONTENDEnv - FRONTENDサーバに関する設定を読み込む
func ReadFRONTENDEnv() (FRONTENDConfig, error) {
	var FRONTENDCfg FRONTENDConfig
	err := envconfig.Process("FRONTEND", &FRONTENDCfg)
	if err != nil {
		return FRONTENDConfig{}, err
	}
	if FRONTENDCfg.Host == "" {
		FRONTENDCfg.Host = "127.0.0.1"
	}
	if FRONTENDCfg.Port == "" {
		FRONTENDCfg.Port = "80"
	}
	return FRONTENDCfg, nil
}
