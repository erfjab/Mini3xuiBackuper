package main

import (
	"Mini3xuiBackuper/internal/client"
	"Mini3xuiBackuper/internal/config"
	"Mini3xuiBackuper/internal/utils"
	"log"
	"os"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	c, err := client.New(cfg.PanelHost)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	if err := c.Login(cfg.PanelUsername, cfg.PanelPassword); err != nil {
		log.Fatalf("Login failed: %v", err)
	}
	log.Println("Logged in successfully")

	data, err := c.DownloadDB()
	if err != nil {
		log.Fatalf("Failed to download database: %v", err)
	}
	log.Printf("Database downloaded (%d bytes)", len(data))

	filename := utils.BackupFilename(cfg.PanelUsername)
	if err := os.WriteFile(filename, data, 0600); err != nil {
		log.Fatalf("Failed to save backup: %v", err)
	}
	log.Printf("Backup saved: %s", filename)
}
