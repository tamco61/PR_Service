package models

// todo проверка по условию корректности моделей

type User struct {
	ID       string `gorm:"primaryKey"`
	Name     string `gorm:"not null"`
	TeamName string `gorm:"default:null"`
	IsActive bool   `gorm:"default:true"`
}
