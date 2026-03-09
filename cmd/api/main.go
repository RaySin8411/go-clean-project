package main

import (
	"go-clean-project/internal/config"
	"log"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// 這裡先只把 config 印出，後續會使用到
	log.Println(config)
}
