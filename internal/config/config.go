package config

import (
	"errors"
	"github.com/ilyakaznacheev/cleanenv"
	"time"
)

var e = errors.New("error while reading cfg")

type Config struct {
	Env        string `yaml:"env" envDefault:"local"`
	ServerHTTP `yaml:"serverHTTP"`
	PgUrl      string `yaml:"pgUrl" env-required:"true"`
	//RedisDB    Redis  `yaml:"redis"`
}

type ServerHTTP struct {
	Addr           string        `yaml:"addr" env-default:"localhost:8080"`
	Timeout        time.Duration `yaml:"timeout" env-default:"4s"`
	ShutDownTimout time.Duration `yaml:"idleTimeout" env-default:"10s"`
}

/*type PG struct {
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	DataBaseName string `yaml:"databaseName"`
	SslMode      string `yaml:"sslMode"`
}*/

func NewConfig(path string) (*Config, error) {
	cfg := &Config{}

	if err := cleanenv.ReadConfig(path, cfg); err != nil {
		return nil, e
	}

	return cfg, nil
}
