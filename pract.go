package main

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"log"
)

type Config struct {
	User string `env:"USERNAME,required"`
}

func main() {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Current user is %s\n", cfg.User)
}
