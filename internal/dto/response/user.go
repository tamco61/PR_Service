package response

import "app/internal/models"

type UserResponse struct {
	ID       string `json:"user_id"`
	Username string `json:"username"`
	TeamName string `json:"team_name"`
	IsActive bool   `json:"is_active"`
}

func ToUserResponse(user *models.User) UserResponse {
	return UserResponse{
		ID:       user.ID,
		Username: user.Name,
		TeamName: user.TeamName,
		IsActive: user.IsActive,
	}
}
