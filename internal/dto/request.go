package request

import "app/internal/models"

type TeamGetReq struct {
	TeamName string `json:"team_name" validate:"required"`
}

type TeamAddReq struct {
	TeamName string        `json:"name" validate:"required"`
	Members  []models.User `json:"members" validate:"required"`
}

type UserGetReviewReq struct {
	ID string `json:"user_id" validate:"required"`
}

type UserSetIsActiveReq struct {
	ID       string `json:"id" validate:"required"`
	IsActive bool   `json:"is_active"`
}

type PullRequestCreateReq struct {
	Name     string `json:"name" validate:"required"`
	AuthorID string `json:"author_id" validate:"required"`
}

type PullRequestMergeReq struct {
	ID string `json:"id" validate:"required"`
}

type PullRequestReassignReq struct {
	PullReqID string `json:"pr_id" validate:"required"`
	OldUserID string `json:"old_user_id" validate:"required"`
}
