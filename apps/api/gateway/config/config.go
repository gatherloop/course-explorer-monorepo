package config

import (
	"course-explorer-monorepo/libs/api/utils/config"
)

type Config struct {
	Server config.ServerConfig
	Api    config.APIConfig
}

func Get() Config {
	return Config{
		Server: config.GetServerConfig(),
		Api:    config.GetAPIConfig(),
	}
}
