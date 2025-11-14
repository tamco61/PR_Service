package handler

import (
	"app/internal/service"

	"github.com/gin-gonic/gin"
)

type TeamHandler struct {
	service *service.TeamService
}

func NewTeamHandler(s *service.TeamService) *TeamHandler {
	return &TeamHandler{service: s}
}

func (h *TeamHandler) Get(c *gin.Context) {

}

func (h *TeamHandler) Add(c *gin.Context) {

}
