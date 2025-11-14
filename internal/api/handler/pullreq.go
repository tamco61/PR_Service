package handler

import (
	"app/internal/service"

	"github.com/gin-gonic/gin"
)

type PullRequestHandler struct {
	service *service.PullRequestService
}

func NewPullRequestHandler(s *service.PullRequestService) *PullRequestHandler {
	return &PullRequestHandler{service: s}
}

func (h *PullRequestHandler) Create(c *gin.Context) {

}

func (h *PullRequestHandler) Merge(c *gin.Context) {

}

func (h *PullRequestHandler) Reassign(c *gin.Context) {

}
