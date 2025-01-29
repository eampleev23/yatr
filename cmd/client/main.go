package main

import (
	"bytes"
	"fmt"
	"github.com/eampleev23/yatr/internal/client_config"
	"log"
	"net/http"
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

	jsonDataStr := `{"queue": "ECODEVTEST", "summary": "Test Issue", "parent":"ECODEVTEST-1968", "type": "milestone", "assignee": "em.ampleev@svo.air.loc"}`
	jsonData := []byte(jsonDataStr)

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Ошибка формирования запроса, попробуйте обновить клиент")
		return fmt.Errorf("http.NewRequest failed %w", err)
	}
	fmt.Println("request", request)
	return nil
}
