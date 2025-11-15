package models

type User struct {
	ID       string `gorm:"primaryKey"`
	Name     string `gorm:"not null"`
	IsActive bool   `gorm:"default:true"`
}
