package main

import (
	"fmt"
	"github.com/eampleev23/yatr/internal/client_config"
	"log"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	c := client_config.NewConfig()
	fmt.Println(c)
	//if err := api_requests.Create(c); err != nil {
	//	return fmt.Errorf("api_requests.Create failed: %w", err)
	//}
	return nil
}
