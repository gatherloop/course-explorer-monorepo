package config

import "os"

type APIConfig struct {
	ContentAPIURL string
}

func GetAPIConfig() APIConfig {
	return APIConfig{
		ContentAPIURL: os.Getenv("CONTENT_API_URL"),
	}
}
