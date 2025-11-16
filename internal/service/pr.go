package service

import (
	"app/internal/models"
	"fmt"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

type PullRequest struct {
	db *gorm.DB
}

func NewPullRequest(db *gorm.DB) *PullRequest {
	return &PullRequest{db: db}
}

func (s *PullRequest) Create(pr_name string, author_id string) (models.PullRequest, error) {
	author := models.User{ID: author_id}
	if err := s.db.First(&author).Error; err != nil {
		return models.PullRequest{}, err
	}

	pullreq := models.PullRequest{
		Name:     pr_name,
		AuthorID: author_id,
		Status:   models.PRStatusOpen,
	}

	tx := s.db.Begin()
	if tx.Error != nil {
		return models.PullRequest{}, tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(&pullreq).Error; err != nil {
		tx.Rollback()
		return models.PullRequest{}, fmt.Errorf("pr exist")
	}

	var team models.Team
	err := tx.
		Preload("Members", "is_active = ?", true).
		Joins("JOIN team_users ON team_users.team_team_name = teams.team_name").
		Where("team_users.user_id = ?", author_id).
		First(&team).Error

	if err != nil {
		tx.Rollback()
		return models.PullRequest{}, err
	}

	var reviewers []models.User

	for _, member := range team.Members {
		if member.ID != author_id {
			reviewers = append(reviewers, member)
		}
	}

	if len(reviewers) == 0 {
		return pullreq, tx.Commit().Error
	}

	count := 2
	if len(reviewers) > count {

		rand.Shuffle(len(reviewers), func(i, j int) {
			reviewers[i], reviewers[j] = reviewers[j], reviewers[i]
		})

		reviewers = reviewers[:count]
	}

	if err := tx.Model(&pullreq).Association("Reviewers").Append(reviewers); err != nil {
		tx.Rollback()
		return models.PullRequest{}, err
	}

	if err := tx.Commit().Error; err != nil {
		return models.PullRequest{}, err
	}

	err = s.db.Preload("Reviewers").First(&pullreq).Error

	return pullreq, err
}

func (s *PullRequest) Merge(pr_id string) (models.PullRequest, error) {
	pullreq := models.PullRequest{ID: pr_id}

	tx := s.db.Begin()
	defer tx.Rollback()
	if err := s.db.First(&pullreq).Error; err != nil {
		return models.PullRequest{}, fmt.Errorf("pr not found")
	}

	if pullreq.Status == models.PRStatusMerge {
		tx.Commit()
		return pullreq, fmt.Errorf("pr merge")
	}

	if pullreq.Status != models.PRStatusOpen {
		tx.Commit()
		return pullreq, fmt.Errorf("pr conflict")
	}

	now := time.Now()
	if err := tx.Model(&pullreq).Updates(map[string]interface{}{
		"status":    models.PRStatusMerge,
		"merged_at": &now,
	}).Error; err != nil {
		return pullreq, err
	}

	pullreq.Status = models.PRStatusMerge
	pullreq.MergedAt = &now

	return pullreq, nil
}

func (s *PullRequest) Reassign(pr_id string, old_user_id string) (models.PullRequest, error) {
	pullreq := models.PullRequest{ID: pr_id}

	tx := s.db.Begin()
	defer tx.Rollback()

	if err := tx.Preload("Reviewers").First(&pullreq).Error; err != nil {
		return pullreq, fmt.Errorf("pr not found")
	}

	if pullreq.Status == models.PRStatusMerge {
		return pullreq, fmt.Errorf("pr already merge")
	}

	oldReviewerId := -1
	for i, reviewer := range pullreq.Reviewers {
		if reviewer.ID == old_user_id {
			oldReviewerId = i
			break
		}
	}

	if oldReviewerId == -1 {
		return pullreq, fmt.Errorf("no candidate")
	}

	var team models.Team
	if err := tx.
		Preload("Members", "is_active = ?", true).
		Joins("JOIN team_users ON team_users.team_team_name = teams.team_name").
		Where("team_users.user_id = ?", pullreq.AuthorID).
		First(&team).Error; err != nil {
		return pullreq, err
	}

	var reviewers []models.User
	for _, m := range team.Members {
		if m.ID == pullreq.AuthorID || m.ID == old_user_id {
			reviewers = append(reviewers, m)
		}
	}

	if len(reviewers) == 0 {
		return pullreq, fmt.Errorf("not assign")
	}

	newReviewer := reviewers[rand.Intn(len(reviewers))]
	if err := tx.Model(&pullreq).Association("Reviewers").Replace(&pullreq.Reviewers[oldReviewerId], &newReviewer); err != nil {
		return pullreq, fmt.Errorf("not assign")
	}

	tx.Commit()

	return pullreq, nil

}
