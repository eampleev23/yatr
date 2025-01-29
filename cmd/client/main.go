package main

import (
	"fmt"
	"github.com/eampleev23/yatr/internal/client_config"
	"log"
	url2 "net/url"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	c := client_config.NewConfig()
	url, err := url2.JoinPath("https://api.tracker.yandex.net/", "v2/issues/")
	if err != nil {
		return fmt.Errorf("url2.JoinPath failed %w", err)
	}
	fmt.Println("c", c)
	fmt.Println("url:", url)
	return nil
}
