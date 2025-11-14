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

	db, err := db.InitDB(cfg.DBPath)
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	userService := service.NewUserService(db)
	teamService := service.NewTeamService(db)
	pullrequestService := service.NewPullRequestService(db)

	userHandler := handler.NewUserHandler(userService)
	teamHandler := handler.NewTeamHandler(teamService)
	pullrequestHandler := handler.NewPullRequestHandler(pullrequestService)

	r := gin.Default()
	routes.RegisterRoutes(r, teamHandler, userHandler, pullrequestHandler)

	if err := r.Run(cfg.ServerAddr); err != nil {
		log.Fatal("failed to run server:", err)
	}

}
