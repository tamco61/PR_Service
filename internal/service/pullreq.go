package service

import "gorm.io/gorm"

type PullRequestService struct {
	db *gorm.DB
}

func NewPullRequestService(db *gorm.DB) *PullRequestService {
	return &PullRequestService{db: db}
}

func (s *PullRequestService) Create() {

}

func (s *PullRequestService) Merge() {

}

func (s *PullRequestService) Reassign() {

}
