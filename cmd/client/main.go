package main

import (
	"fmt"
	"github.com/eampleev23/yatr/internal/client_config"
	"github.com/eampleev23/yatr/internal/services"
	"log"
	"unicode/utf8"
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
	}
	return nil
}

func trimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}
