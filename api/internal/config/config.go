package config

import "os"

type Config struct {
	AuthAddr string
	Port     string
}

func Load() Config {
	return Config{
		AuthAddr: getEnv("AUTH_GRPC_ADDR", "localhost:50051"),
		Port:     getEnv("PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
