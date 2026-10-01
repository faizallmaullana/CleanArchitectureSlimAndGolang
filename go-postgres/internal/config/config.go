package config

import "os"

type Config struct {
	HTTPAddr string
	Database DatabaseConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

func Load() Config {
	return Config{
		HTTPAddr: getEnv("HTTP_ADDR", ":8080"),
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "app_user"),
			Password: getEnv("DB_PASSWORD", "app_password"),
			Name:     getEnv("DB_NAME", "app_db"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
	}
}

func (config DatabaseConfig) DSN() string {
	return "postgres://" + config.User + ":" + config.Password + "@" + config.Host + ":" + config.Port + "/" + config.Name + "?sslmode=" + config.SSLMode
}

func getEnv(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}
