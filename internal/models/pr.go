package models

import "time"

type PRStatus string

const (
	PRStatusOpen   PRStatus = "OPEN"
	PRStatusMerge  PRStatus = "MERGE"
	PRStatusClosed PRStatus = "CLOSED"
)

type PullRequest struct {
	ID        string   `gorm:"primaryKey"`
	Name      string   `gorm:"not null"`
	AuthorID  string   `gorm:"not null"`
	Status    PRStatus `gorm:"default:'OPEN'"`
	Reviewers []User   `gorm:"many2many:reviewers"`

	MergedAt *time.Time `gorm:"default:null"`
}
