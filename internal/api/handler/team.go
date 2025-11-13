package handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TeamHandler struct {
	db *gorm.DB
}

func NewTeamHandler(db *gorm.DB) *TeamHandler {
	return &TeamHandler{db: db}
}

func (h *TeamHandler) Get(c *gin.Context) {

}

func (h *TeamHandler) Add(c *gin.Context) {

}
