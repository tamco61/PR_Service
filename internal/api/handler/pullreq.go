package handler

import "gorm.io/gorm"

type PullRequestHandler struct {
	db *gorm.DB
}

func NewPullRequestHandler(db *gorm.DB) *PullRequestHandler {
	return &PullRequestHandler{db: db}
}
