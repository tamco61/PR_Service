package handler

import (
	"app/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) GetReview(c *gin.Context) {

}

func (h *UserHandler) SetIsActive(c *gin.Context) {

}
