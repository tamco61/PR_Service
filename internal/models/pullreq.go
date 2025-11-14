package models

type PullRequest struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	AuthorID  uint   `gorm:"not null"`
	Status    bool   `gorm:"default:false"`
	Reviewers []User `gorm:"many2many:reviewers"`
}
