package env

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/tetsuzawa/Goonstone/containers/api/pkg/redis"
)

// ReadRedisEnv - 指定したenvfileからRedisに関する設定を読み込む
func ReadRedisEnv() (redis.Config, error) {
	var RedisCfg redis.Config
	err := envconfig.Process("REDIS", &RedisCfg)
	if err != nil {
		return redis.Config{}, err
	}
	return RedisCfg, nil
}
