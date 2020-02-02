package utils

import (
	"github.com/sirupsen/logrus"
	"github.com/tomazk/envcfg"
)

type AppConfig struct {
	PROD_MODE  string
	APP_PORT   int
	DB_USER    string
	DB_PASS    string
	DB_HOST    string
	DB_PORT    int
	DB_NAME    string
	JWT_SECRET string
	DOMAIN     string
}

var Config AppConfig

func NewConfig() {
	err := envcfg.Unmarshal(&Config)
	if err != nil {
		logrus.Warn("Failed to fill config from environment variables")
	}
	if Config.JWT_SECRET == "" {
		logrus.Warn("JWT Secret not set. Falling back to default. This is insecure. Fix ASAP.")
		Config.JWT_SECRET = "XL$jBmE6jwkC67i=9TgF"
	}
	if Config.DB_HOST == "" {
		Config.DB_HOST = "localhost"
	}
	if Config.DB_PORT == 0 {
		Config.DB_PORT = 3306
	}
	if Config.DB_NAME == "" {
		Config.DB_NAME = "labonte"
	}
	if Config.DB_USER == "" {
		Config.DB_USER = "root"
	}
	if Config.DOMAIN == "" {
		Config.DOMAIN = "localhost:8080"
	}
}
