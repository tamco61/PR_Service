package handler

// todo прописать ошибки и статус коды
import (
	"app/internal/dto"
	requests "app/internal/dto"
	"app/internal/service"
	"app/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PullRequest struct {
	service *service.PullRequest
}

func NewPullRequest(s *service.PullRequest) *PullRequest {
	return &PullRequest{service: s}
}

func (h *PullRequest) Create(c *gin.Context) {
	var req requests.PullRequestCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pr, err := h.service.Create(req.Name, req.AuthorID)
	if err != nil {
		if err.Error() == "pr exist" {
			c.JSON(http.StatusBadRequest, utils.JSONError(dto.ErrorCodePRExists, "PR id already exists"))
		} else {
			c.JSON(http.StatusConflict, utils.JSONError(dto.ErrorCodeNotFound, "author/team not found"))
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, pr)
}

func (h *PullRequest) Merge(c *gin.Context) {
	var req requests.PullRequestMergeReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pr, err := h.service.Merge(req.ID)
	if err != nil {
		if err.Error() == "pr not found" {
			c.JSON(http.StatusNotFound, utils.JSONError(dto.ErrorCodeNotFound, "pr not found"))
		} else {
			c.JSON(http.StatusConflict, utils.JSONError(dto.ErrorCodeNotFound, "pr conflict"))
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, pr)
}

func (h *PullRequest) Reassign(c *gin.Context) {
	var req requests.PullRequestReassignReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pr, err := h.service.Reassign(req.PullReqID, req.OldUserID)
	if err != nil {
		if err.Error() == "pr not found" {
			c.JSON(http.StatusBadRequest, utils.JSONError(dto.ErrorCodeNotFound, "pr/user not found"))
		} else if err.Error() == "pr already merge" {
			c.JSON(http.StatusConflict, utils.JSONError(dto.ErrorCodeNotAssigned, "cannot reassign on merged PR"))
		} else if err.Error() == "no candidate" {
			c.JSON(http.StatusConflict, utils.JSONError(dto.ErrorCodeNoCandidate, "no active replacement candidate in team"))
		} else {
			c.JSON(http.StatusConflict, utils.JSONError(dto.ErrorCodeNotFound, "reviewer is not assigned to this PR"))
		}
		return
	}

	c.JSON(http.StatusOK, pr)
}
