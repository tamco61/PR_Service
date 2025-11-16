package response

import "app/internal/models"

type TeamMemberDTO struct {
	ID       string `json:"user_id"`
	Username string `json:"username"`
	IsActive bool   `json:"is_active"`
}

type TeamResponse struct {
	TeamName string          `json:"team_name"`
	Members  []TeamMemberDTO `json:"members"`
}

func ToTeamMemberResponse(users []models.User) []TeamMemberDTO {
	result := make([]TeamMemberDTO, len(users))
	for i, u := range users {
		result[i] = TeamMemberDTO{
			ID:       u.ID,
			Username: u.Name,
			IsActive: u.IsActive,
		}
	}
	return result
}

func ToTeamResponse(team *models.Team) TeamResponse {
	return TeamResponse{
		TeamName: team.TeamName,
		Members:  ToTeamMemberResponse(team.Members),
	}
}
