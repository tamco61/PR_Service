package handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	db *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{db: db}
}

func (h *UserHandler) GetReview(c *gin.Context) {

}

func (h *UserHandler) SetIsActive(c *gin.Context) {

}
