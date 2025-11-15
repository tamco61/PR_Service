package service

import (
	"app/internal/models"

	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) GetReview(id uint) ([]models.PullRequest, error) {
	// var reviews []models.PullRequest
	// err := s.db.Where("AuthorID")
}

func (s *UserService) SetIsActive() {

}
