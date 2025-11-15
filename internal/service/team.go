package service

import (
	"app/internal/models"

	"gorm.io/gorm"
)

type Team struct {
	db *gorm.DB
}

func NewTeam(db *gorm.DB) *Team {
	return &Team{db: db}
}

func (s *Team) Get(team_name string) (models.Team, error) {
	team := models.Team{TeamName: team_name}
	err := s.db.First(&team).Error

	return team, err
}

func (s *Team) Add(team_name string, members []models.User) (models.Team, error) {
	team := models.Team{TeamName: team_name, Members: members}
	err := s.db.Save(team).Error

	return team, err
}
