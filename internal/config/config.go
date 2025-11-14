package config

import (
	"os"
)

type Config struct {
	DBDSN      string
	ServerAddr string
}

func Load() *Config {
	return &Config{
		DBDSN:      getEnv("DB_DSN", "host=localhost user=appuser password=secretpassword dbname=appdb port=5432 sslmode=disable"),
		ServerAddr: getEnv("SERVER_ADDR", ":8080"),
	}
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}

	return fallback
}
