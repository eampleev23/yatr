package client_config

import (
	"flag"
	"net/http"
	"net/http/cookiejar"
	"os"
)

type Config struct {
	YTrToken   string
	IsCreating bool
	FilePath   string
	CloudOrgId string
	HttpClient *http.Client
}

func NewConfig() *Config {
	config := &Config{}
	config.SetValues()
	jar, _ := cookiejar.New(nil)
	config.HttpClient = &http.Client{
		Jar: jar,
	}
	return config
}

func (c *Config) SetValues() {
	flag.StringVar(&c.YTrToken, "token", "", "Set yandex tracker token")
	flag.BoolVar(&c.IsCreating, "create", false, "Is creating new tasks or updating")
	flag.StringVar(&c.FilePath, "file", "", "File path")
	flag.StringVar(&c.CloudOrgId, "oid", "", "X-Cloud-Org-Id Header")
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

	if envCloudOrgId := os.Getenv("CLOUD_ORG_ID"); envCloudOrgId != "" {
		c.CloudOrgId = envCloudOrgId
	}
}
