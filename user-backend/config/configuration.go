package config

import (
	"os"
)

const (
	LOCAL       = "local"
	DEVELOPMENT = "development"
)

// ENVIRONMENT:
const ENVIRONMENT string = LOCAL // LOCAL, DEVELOPMENT, PRODUCTION

var env = map[string]map[string]string{
	// local environment configuration
	"local": {
		"APP_NAME" : "fse",
		"SECRET_KEY" : "secret",

	},
}

// CONFIG : global configuration
var CONFIG = env[ENVIRONMENT]

// Getenv : function for Environment Lookup
func Getenv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func InitConfig() {
	for key := range CONFIG {
		CONFIG[key] = Getenv(key, CONFIG[key])
		os.Setenv(key, CONFIG[key])
	}
}
