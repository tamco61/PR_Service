package requests

type TeamGetReq struct {
	ID uint `json:"id" validate:"required"`
}

type TeamAddReq struct {
	Name string `json:"name" validate:"required"`
}

type UserGetReviewReq struct {
	ID uint `json:"id" validate:"required"`
}

type UserSetIsActiveReq struct {
	ID       uint `json:"id" validate:"required"`
	IsActive bool `json:"is_active"`
}

type PullRequestCreateReq struct {
	Name        string `json:"name" validate:"required"`
	AuthorID    uint   `json:"author_id" validate:"required"`
	ReviewersID []uint `json:"reviewers_id" validate:"required"`
}

type PullRequestMergeReq struct {
	ID uint `json:"id" validate:"required"`
}

type PullRequestReassignReq struct {
	PullReqID uint `json:"pr_id" validate:"required"`
	AuthorID  uint `json:"author_id" validate:"required"`
}
