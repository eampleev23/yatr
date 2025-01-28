package main

import "log"

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	c, err := client_config.NewConfig()
}
