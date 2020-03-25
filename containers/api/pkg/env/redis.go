package env

import (
	"github.com/kelseyhightower/envconfig"

	"github.com/tetsuzawa/Goonstone/containers/api/pkg/redisx"
)

// ReadRedisEnv - 指定したenvfileからRedisに関する設定を読み込む
func ReadRedisEnv() (redisx.Config, error) {
	var RedisCfg redisx.Config
	err := envconfig.Process("REDIS", &RedisCfg)
	if err != nil {
		return redisx.Config{}, err
	}
	return RedisCfg, nil
}
