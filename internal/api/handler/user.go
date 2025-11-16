package handler

import (
	"app/internal/dto"
	requests "app/internal/dto"
	"app/internal/service"
	"app/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	service *service.User
}

func NewUser(s *service.User) *User {
	return &User{service: s}
}

func (h *User) GetReview(c *gin.Context) {
	var req requests.UserGetReviewReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.GetReview(req.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)

}

func (h *User) SetIsActive(c *gin.Context) {
	var req requests.UserSetIsActiveReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.SetIsActive(req.ID, req.IsActive)
	if err != nil {
		c.JSON(http.StatusNotFound, utils.JSONError(dto.ErrorCodeNotFound, "user not found"))
		return
	}

	c.JSON(http.StatusOK, user)
}
