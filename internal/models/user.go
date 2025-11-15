package models

// todo проверка по условию корректности моделей

type User struct {
	ID       string `gorm:"primaryKey"`
	Name     string `gorm:"not null"`
	IsActive bool   `gorm:"default:true"`
}
