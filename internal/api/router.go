package routes

import (
	"app/internal/api/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, teamHandler *handler.TeamHandler, userHandler *handler.UserHandler, pullrequestHandler *handler.PullRequestHandler) {
	team := r.Group("/team")
	{
		team.GET("/team/get", teamHandler.Get)
		team.POST("/team/add", teamHandler.Add)
	}

	users := r.Group("/users")
	{
		users.GET("/users/getReview", userHandler.GetReview)
		users.POST("/users/setIsActive", userHandler.SetIsActive)
	}

	pullrequest := r.Group("/pullRequest")
	{
		pullrequest.POST("/pullRequest/create", pullrequestHandler.Create)
		pullrequest.POST("/pullRequest/merge", pullrequestHandler.Merge)
		pullrequest.POST("/pullRequest/reassign", pullrequestHandler.Reassign)
	}
}
