package dto

type TeamGetRequest struct {
	ID uint `json:"id"  validate:"required"`
}

type TeamAddRequest struct {
	Name string `json:"name" validate:"required"`
}

type UserGetReviewRequest struct {
	ID uint `json:"id" validate:"required"`
}

type UserSetIsActiveRequest struct {
	ID       uint `json:"id" validate:"required"`
	IsActive bool `json:"is_active"`
}

type PullRequestCreateRequest struct {
	Name        string `json:"name" validate:"required"`
	AuthorID    uint   `json:"author_id" validate:"required"`
	ReviewersID []uint `json:"reviewers_id" validate:"required"`
}

type PullRequestMergeRequest struct {
	ID uint `json:"id" validate:"required"`
}

type PullRequestReassignRequest struct {
	PullReqID uint `json:"pullreq_id" validate:"required"`
	AuthorID  uint `json:"author_id" validate:"required"`
}
