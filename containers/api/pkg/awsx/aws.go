package awsx

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/kelseyhightower/envconfig"
)

// Config - AWSの接続情報に関する設定
type Config struct {
	Profile  string `split_words:"true"`
	S3Bucket string `split_words:"true"`
}

type Connection struct {
	Config  Config
	Session *session.Session
	SVC     *s3.S3
}

// ReadEnv - 指定したenvfileからAWSに関する設定を読み込む
func ReadEnv(cfg *Config) error {
	err := envconfig.Process("AWS", cfg)
	if err != nil {
		return fmt.Errorf("envconfig.Process: %w", err)
	}
	return nil
}

func (c Config) build() Config {
	if c.Profile == "" {
		c.Profile = session.DefaultSharedConfigProfile
	}
	return c
}

// Connect - AWSに接続
func Connect(c Config) (*Connection, error) {
	c = c.build()
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile:           c.Profile,
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		return nil, fmt.Errorf("session.NewSessionWithOptions: %w", err)
	}
	svc := s3.New(sess)
	return &Connection{Config: c, Session: sess, SVC:svc}, nil
}
