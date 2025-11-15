package routes

import (
	"app/internal/api/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, Team *handler.Team, User *handler.User, PullRequest *handler.PullRequest) {
	team := r.Group("/team")
	{
		team.GET("/get", Team.Get)
		team.POST("/add", Team.Add)
	}

	users := r.Group("/users")
	{
		users.GET("/getReview", User.GetReview)
		users.POST("/setIsActive", User.SetIsActive)
	}

	pullrequest := r.Group("/pullRequest")
	{
		pullrequest.POST("/create", PullRequest.Create)
		pullrequest.POST("/merge", PullRequest.Merge)
		pullrequest.POST("/reassign", PullRequest.Reassign)
	}
}
