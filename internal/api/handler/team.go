package handler

import (
	requests "app/internal/dto"
	"app/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TeamHandler struct {
	service *service.TeamService
}

func NewTeamHandler(s *service.TeamService) *TeamHandler {
	return &TeamHandler{service: s}
}

func (h *TeamHandler) Get(c *gin.Context) {
	var req requests.TeamGetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func (h *TeamHandler) Add(c *gin.Context) {
	var req requests.TeamAddRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

}
