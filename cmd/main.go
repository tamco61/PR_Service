package main

import (
	routes "app/internal/api"
	"app/internal/api/handler"
	"app/internal/config"
	"app/internal/db"
	"app/internal/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	db, err := db.InitDB(cfg.DBDSN)
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	userService := service.NewUser(db)
	teamService := service.NewTeam(db)
	pullRequestService := service.NewPullRequest(db)

	userHandler := handler.NewUser(userService)
	teamHandler := handler.NewTeam(teamService)
	pullRequestHandler := handler.NewPullRequest(pullRequestService)

	r := gin.Default()
	routes.RegisterRoutes(r, teamHandler, userHandler, pullRequestHandler)

	if err := r.Run(cfg.ServerAddr); err != nil {
		log.Fatal("failed to run server:", err)
	}

}
