package client_config

import (
	"flag"
	"os"
)

type Config struct {
	YTrToken   string
	IsCreating bool
	FilePath   string
}

func NewConfig() *Config {
	config := &Config{}
	config.SetValues()
	return config
}

func (c *Config) SetValues() {
	flag.StringVar(&c.YTrToken, "token", "", "Set yandex tracker token")
	flag.BoolVar(&c.IsCreating, "create", false, "Is creating new tasks or updating")
	flag.StringVar(&c.FilePath, "file", "", "File path")
	flag.Parse()

	if envYTrToken := os.Getenv("YTR_TOKEN"); envYTrToken != "" {
		c.YTrToken = envYTrToken
	}

	if envIsCreating := os.Getenv("IS_CREATING"); envIsCreating == "true" {
		c.IsCreating = true
	}

	if envFilePath := os.Getenv("FILE_PATH"); envFilePath != "" {
		c.FilePath = envFilePath
	}
}
