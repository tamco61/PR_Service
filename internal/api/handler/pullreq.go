package handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PullRequestHandler struct {
	db *gorm.DB
}

func NewPullRequestHandler(db *gorm.DB) *PullRequestHandler {
	return &PullRequestHandler{db: db}
}

func (h *PullRequestHandler) Create(c *gin.Context) {

}

func (h *PullRequestHandler) Merge(c *gin.Context) {

}

func (h *PullRequestHandler) Reassign(c *gin.Context) {

}
