package models

// todo проверка по условию корректности моделей

type Team struct {
	TeamName string `gorm:"primaryKey"`
	Members  []User `gorm:"many2many:team_users"`
}
