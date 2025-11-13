package routes

import (
	"app/internal/api/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, teamHandler *handler.TeamHandler, userHandler *handler.UserHandler, pullrequestHandler *handler.PullRequestHandler) {
	team := r.Group("/team")
	{
		team.GET("/get", teamHandler.Get)
		team.POST("/add", teamHandler.Add)
	}

	users := r.Group("/users")
	{
		users.GET("/getReview", userHandler.GetReview)
		users.POST("/setIsActive", userHandler.SetIsActive)
	}

	pullrequest := r.Group("/pullRequest")
	{
		pullrequest.POST("/create", pullrequestHandler.Create)
		pullrequest.POST("/merge", pullrequestHandler.Merge)
		pullrequest.POST("/reassign", pullrequestHandler.Reassign)
	}
}
