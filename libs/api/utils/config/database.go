package config

import "os"

type DatabaseConfig struct {
	Host     string
	Username string
	Password string
	Name     string
	Port     string
}

func GetDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		Host:     os.Getenv("DB_HOST"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
	}
}
