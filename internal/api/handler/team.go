package handler

import (
	"app/internal/dto"
	requests "app/internal/dto"
	"app/internal/service"
	"app/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Team struct {
	service *service.Team
}

func NewTeam(s *service.Team) *Team {
	return &Team{service: s}
}

func (h *Team) Get(c *gin.Context) {
	var req requests.TeamGetReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	team, err := h.service.Get(req.TeamName)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.JSONError(dto.ErrorCodeNotFound, "team not found"))
		return
	}

	c.JSON(http.StatusOK, team)
}

func (h *Team) Add(c *gin.Context) {
	var req requests.TeamAddReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	team, err := h.service.Add(req.TeamName, req.Members)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.JSONError(dto.ErrorCodeTeamExists, "team_name already exists"))
		return
	}

	c.JSON(http.StatusCreated, team)

}
