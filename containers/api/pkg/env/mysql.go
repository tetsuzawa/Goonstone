package env

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/tetsuzawa/Goonstone/containers/api/pkg/mysql"
)

// ReadMysqlEnv - 指定したenvfileからMysqlに関する設定を読み込む
func ReadMysqlEnv() (mysql.Config, error) {
	var MysqlCfg mysql.Config
	err := envconfig.Process("MYSQL", &MysqlCfg)
	if err != nil {
		return mysql.Config{}, err
	}
	return MysqlCfg, nil
}
