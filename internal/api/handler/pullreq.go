package handler

import (
	requests "app/internal/dto"
	"app/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PullRequestHandler struct {
	service *service.PullRequestService
}

func NewPullRequestHandler(s *service.PullRequestService) *PullRequestHandler {
	return &PullRequestHandler{service: s}
}

func (h *PullRequestHandler) Create(c *gin.Context) {
	var req requests.PullRequestCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func (h *PullRequestHandler) Merge(c *gin.Context) {
	var req requests.PullRequestMergeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func (h *PullRequestHandler) Reassign(c *gin.Context) {
	var req requests.PullRequestReassignRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
