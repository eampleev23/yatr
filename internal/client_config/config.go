package client_config

import "flag"

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
	flag.StringVar(&c.YTrToken, "yt", "_", "Set ytracker token")
	flag.BoolVar(&c.IsCreating, "is_c", false, "Is creating new tasks or updating")
	flag.StringVar(&c.FilePath, "file", "", "File path")
}
