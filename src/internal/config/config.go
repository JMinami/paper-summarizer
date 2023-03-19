package config

import (
	"os"
)

type Config struct {
	LineMessagingAPIChannelAccessToken string
	LineMessagingAPIChannelSecret      string
}

func New() *Config {
	return &Config{
		LineMessagingAPIChannelAccessToken: getEnv("LINE_MESSAGING_API_CHANNEL_ACCESS_TOKEN", ""),
		LineMessagingAPIChannelSecret:      getEnv("LINE_MESSAGING_API_CHANNEL_SECRET", ""),
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}