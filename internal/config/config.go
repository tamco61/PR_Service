package config

import (
	"os"
)

type Config struct {
	DBPath     string
	ServerAddr string
}

func Load() *Config {
	return &Config{
		DBPath:     getEnv("SQLITE_DB_PATH", "database.db"),
		ServerAddr: getEnv("SERVER_ADDR", ":8080"),
	}
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}

	return fallback
}
