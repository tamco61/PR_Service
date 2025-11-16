package utils

import (
	"app/internal/dto/response"
)

func JSONError(code response.ErrorCode, message string) response.ErrorResponse {
	return response.ErrorResponse{
		Error: response.ErrorDetail{
			Code:    code,
			Message: message,
		},
	}
}
