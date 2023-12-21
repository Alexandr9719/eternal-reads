package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv string
	Port   int
	Token  string
}

var (
	testConfig Config = Config{
		Token: "XXX",
	}
	productionConfig Config = Config{
		Token: "YYY",
	}
)

func NewConfig() *Config {
	return &Config{
		AppEnv: "production",
		Port:   443,
	}
}

func (c *Config) LoadFromEnvFile(filename string) error {
	err := godotenv.Load(filename)

	if err != nil {
		return err
	}

	return nil
}

func (c *Config) LoadFromDefaultEnvFile() error {
	err := godotenv.Load()

	if err != nil {
		return err
	}

	return nil
}

func (c *Config) InitConfigByEnv() {
	if value, exists := os.LookupEnv("APP_ENV"); exists {
		switch value {
		case "production":
			c.AppEnv = "production"
		case "testing":
			c.AppEnv = "testing"
		case "development":
			c.AppEnv = "development"
		default:
			c.AppEnv = "production"
		}
	}

	if value, exists := os.LookupEnv("PORT"); exists {
		port, err := strconv.Atoi(value)

		if err == nil {
			c.Port = port
		}
	}

	switch c.AppEnv {
	case "testing":
	case "development":
		c.Token = testConfig.Token
	case "production":
		c.Token = productionConfig.Token
	}
}
