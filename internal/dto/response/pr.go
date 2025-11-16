package response

import (
	"app/internal/models"
	"time"
)

type PullReqResponse struct {
	PRid        string   `json:"pull_request_id"`
	PRname      string   `json:"pull_request_name"`
	AuthorID    string   `json:"author_id"`
	PRstatus    string   `json:"status"`
	PRreviewers []string `json:"assigned_reviewers"`
	PRcreatedAt string   `json:"created_at"`
	PRmergeAt   string   `json:"merge_at"`
}

type PullReqShortResponse struct {
	PRid     string `json:"pull_request_id"`
	PRname   string `json:"pull_request_name"`
	AuthorID string `json:"author_id"`
	Status   string `json:"status"`
}

type ReviewsResponse struct {
	ID           string                 `json:"user_id"`
	PullRequests []PullReqShortResponse `json:"pull_requests"`
}

func ToPullReqResponse(pr *models.PullRequest) PullReqResponse {
	var reviewers []string
	for _, r := range pr.Reviewers {
		reviewers = append(reviewers, r.ID)
	}

	return PullReqResponse{
		PRid:        pr.ID,
		PRname:      pr.Name,
		AuthorID:    pr.AuthorID,
		PRstatus:    string(pr.Status),
		PRreviewers: reviewers,
		PRcreatedAt: pr.CreatedAt.Format(time.RFC3339),
		PRmergeAt:   pr.MergedAt.Format(time.RFC3339),
	}
}

func ToPullReqShortResponse(pr *models.PullRequest) PullReqShortResponse {
	return PullReqShortResponse{
		PRid:     pr.ID,
		PRname:   pr.Name,
		AuthorID: pr.AuthorID,
		Status:   string(pr.Status),
	}
}

func ToReviewsResponse(userId string, prs []models.PullRequest) ReviewsResponse {
	shortPRs := make([]PullReqShortResponse, len(prs))
	for i, pr := range prs {
		shortPRs[i] = ToPullReqShortResponse(&pr)
	}

	return ReviewsResponse{
		ID:           userId,
		PullRequests: shortPRs,
	}
}
