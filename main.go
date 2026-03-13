package main

import (
	"Mini3xuiBackuper/internal/config"
	"log"
)

func main() {
	_, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	log.Printf("Config loaded successfully")
}