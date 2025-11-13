package main

import (
	routes "app/internal/api"
	"app/internal/api/handler"
	"app/internal/config"
	"app/internal/db"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	db, err := db.InitDB(cfg.DBPath)
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	userHandler := handler.NewUserHandler(db)
	teamHandler := handler.NewTeamHandler(db)
	pullrequestHandler := handler.NewPullRequestHandler(db)

	r := gin.Default()
	routes.RegisterRoutes(r, teamHandler, userHandler, pullrequestHandler)

	if err := r.Run(cfg.ServerAddr); err != nil {
		log.Fatal("failed to run server:", err)
	}

}
