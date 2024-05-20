package config

import (
	"time"

	"github.com/joeshaw/envdecode"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

type Config struct {
	ServiceName string `env:"SERVICE_NAME,default=tracerstudy-tracer-service"`
	Port        Port
	MySQL1       MySQL1
	MySQL2       MySQL2
	SIAK_API    SIAK_API
	JWT         JWTConfig
}

type Port struct {
	GRPC string `env:"PORT_GRPC,default=8081"`
	REST string `env:"PORT_REST,default=8080"`
}

type MySQL1 struct {
	Host     string `env:"MYSQL1_HOST"`
	Port     string `env:"MYSQL1_PORT"`
	User     string `env:"MYSQL1_USER"`
	Password string `env:"MYSQL1_PASSWORD"`
	Name     string `env:"MYSQL1_NAME"`
}

type MySQL2 struct {
	Host     string `env:"MYSQL2_HOST"`
	Port     string `env:"MYSQL2_PORT"`
	User     string `env:"MYSQL2_USER"`
	Password string `env:"MYSQL2_PASSWORD"`
	Name     string `env:"MYSQL2_NAME"`
}

type SIAK_API struct {
	URL string `env:"SIAK_API_URL"`
	KEY string `env:"SIAK_API_KEY"`
}

type JWTConfig struct {
	JwtSecretKey  string        `env:"JWT_SECRET_KEY"`
	TokenDuration time.Duration `env:"JWT_DURATION,default=30m"`
}

func NewConfig(env string) (*Config, error) {
	_ = godotenv.Load(env)

	var config Config
	if err := envdecode.Decode(&config); err != nil {
		return nil, errors.Wrap(err, "ERROR: [NewConfig] Error while decoding env")
	}

	return &config, nil
}
