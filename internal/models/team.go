package models

type Team struct {
	TeamName string `gorm:"primaryKey"`
	Members  []User `gorm:"many2many:team_users"`
}
