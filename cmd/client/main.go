package main

import (
	"fmt"
	"github.com/eampleev23/yatr/internal/client_config"
	"github.com/eampleev23/yatr/internal/services"
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
	if c.IsCreating {
		if err := services.GenerateIssues(c); err != nil {
			return fmt.Errorf("could not generate issues: %w", err)
		}
	} else {
		if err := services.UpdateIssues(c); err != nil {
			return fmt.Errorf("could not update issues: %w", err)
		}
	}
	return nil
}
