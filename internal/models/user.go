package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"not null"`
	IsActive bool   `gorm:"default:true"`
}
