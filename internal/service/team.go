package service

import "gorm.io/gorm"

type TeamService struct {
	db *gorm.DB
}

func NewTeamService(db *gorm.DB) *TeamService {
	return &TeamService{db: db}
}

func (s *TeamService) Get() {

}

func (s *TeamService) Add() {

}
