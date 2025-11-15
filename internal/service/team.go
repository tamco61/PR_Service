package service

import (
	"app/internal/models"

	"gorm.io/gorm"
)

type TeamService struct {
	db *gorm.DB
}

func NewTeamService(db *gorm.DB) *TeamService {
	return &TeamService{db: db}
}

func (s *TeamService) Get(team_name string) (models.Team, error) {
	team := models.Team{TeamName: team_name}
	err := s.db.First(&team).Error

	return team, err
}

func (s *TeamService) Add(team_name string, members []models.User) (models.Team, error) {
	team := models.Team{TeamName: team_name, Members: members}
	err := s.db.Save(team).Error

	return team, err
}
