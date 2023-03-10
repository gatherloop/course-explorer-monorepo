package config

import (
	"course-explorer-monorepo/libs/api/utils/config"
)

type Config struct {
  Database config.DatabaseConfig
}

func Get() Config {
	return Config{
    Database: config.GetDatabaseConfig(),
	}
}
