package config

import (
	"os"
	"strconv"
	"sync"
)

type AppEnv string

const (
	AppEnvProduction  AppEnv = "production"
	AppEnvTesting     AppEnv = "testing"
	AppEnvDevelopment AppEnv = "development"
)

type Config struct {
	AppEnv AppEnv `json:"app_env,omitempty"`
	Port   int    `json:"port,omitempty"`
	Token  string `json:"token,omitempty"`
}

var (
	testConfig Config = Config{
		Token: "XXX",
	}
	productionConfig Config = Config{
		Token: "YYY",
	}
	config Config
	once   sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		cfg := initConfigByEnv()

		config = cfg
	})

	return &config
}

func initConfigByEnv() Config {
	var c Config
	if value, exists := os.LookupEnv("APP_ENV"); exists {
		switch value {
		case "production":
			c.AppEnv = AppEnvProduction
		case "testing":
			c.AppEnv = AppEnvTesting
		case "development":
			c.AppEnv = AppEnvDevelopment
		default:
			c.AppEnv = AppEnvProduction
		}
	}

	if value, exists := os.LookupEnv("PORT"); exists {
		port, err := strconv.Atoi(value)

		if err == nil {
			c.Port = port
		} else {
			c.Port = 443
		}
	}

	switch c.AppEnv {
	case AppEnvDevelopment:
	case AppEnvTesting:
		c.Token = testConfig.Token
	case AppEnvProduction:
		c.Token = productionConfig.Token
	}

	return c
}
