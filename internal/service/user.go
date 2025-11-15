package service

import (
	"app/internal/models"

	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{db: db}
}

func (s *User) GetReview(user_id string) ([]models.PullRequest, error) {
	user := models.User{ID: user_id}
	if err := s.db.First(&user).Error; err != nil {
		return nil, err
	}

	var pullreqs []models.PullRequest
	err := s.db.
		Preload("Reviewers").
		Joins("JOIN reviewers ON reviewers.pull_request_id = pull_requests.id").
		Where("reviewers.user_id = ?", user_id).
		Where("pull_requests.status = ?", false).
		Find(&pullreqs).Error

	return pullreqs, err
}

func (s *User) SetIsActive(user_id string, is_active bool) (models.User, error) {
	user := models.User{ID: user_id, IsActive: is_active}
	if err := s.db.First(&user).Error; err != nil {
		return models.User{}, err
	}

	if err := s.db.Model(&user).Update("is_active", is_active).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}
