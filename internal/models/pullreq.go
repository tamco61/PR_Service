package models

type PullRequest struct {
	ID        uint
	Name      string
	Author    User
	Status    bool
	Reviewers []User
}
