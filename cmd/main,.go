package main

import (
	"app/internal/config"
	"app/internal/db"
	"log"
)

func main() {
	cfg := config.Load()

	db, err := db.InitDB(cfg.DBPath)
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

}
