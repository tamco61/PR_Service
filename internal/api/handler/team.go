package handler

import (
	requests "app/internal/dto"
	"app/internal/dto/response"
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
		c.JSON(http.StatusBadRequest, utils.JSONError(response.ErrorCodeNotFound, "team not found"))
		return
	}

	resp := response.ToTeamResponse(&team)
	c.JSON(http.StatusOK, resp)
}

func (h *Team) Add(c *gin.Context) {
	var req requests.TeamAddReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	team, err := h.service.Add(req.TeamName, req.Members)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.JSONError(response.ErrorCodeTeamExists, "team_name already exists"))
		return
	}

	resp := response.ToTeamResponse(&team)
	c.JSON(http.StatusCreated, resp)

}
