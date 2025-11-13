package models

type Team struct {
	ID    uint   `gorm:"primaryKey"`
	Users []User `gorm:"many2many:team_users"`
}
