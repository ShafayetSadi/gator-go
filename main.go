package main

import (
	"fmt"
	"log"

	"github.com/shafayetsadi/gator/internal/config"
)

func main() {
	cfg := config.Config{}
	if err := cfg.Read(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(cfg)

	err := cfg.SetUser("Shafayet Sadi")
	if err != nil {
		log.Fatal("Failed to set current user")
	}
	if err := cfg.Read(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(cfg)
}
