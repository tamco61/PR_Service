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
	for i := range members {
		s.db.Save(members[i])
	}

	team := models.Team{TeamName: team_name, Members: members}
	err := s.db.Save(team).Error

	var userIDs []string
	for _, member := range members {
		userIDs = append(userIDs, member.ID)
	}

	if len(userIDs) == 0 {
		return team, nil
	}

	updateResult := s.db.Model(&models.User{}).
		Where("id IN ?", userIDs).
		Update("team_name", team_name)

	if updateResult.Error != nil {
		return team, updateResult.Error
	}

	return team, err
}
