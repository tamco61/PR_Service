package handler

import (
	requests "app/internal/dto"
	"app/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PullRequest struct {
	service *service.PullRequest
}

func NewPullRequest(s *service.PullRequest) *PullRequest {
	return &PullRequest{service: s}
}

func (h *PullRequest) Create(c *gin.Context) {
	var req requests.PullRequestCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pr, err := h.service.Create(req.Name, req.AuthorID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, pr)
}

func (h *PullRequest) Merge(c *gin.Context) {
	var req requests.PullRequestMergeReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pr, err := h.service.Merge(req.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, pr)
}

func (h *PullRequest) Reassign(c *gin.Context) {
	var req requests.PullRequestReassignReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pr, err := h.service.Reassign(req.PullReqID, req.OldUserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, pr)
}
