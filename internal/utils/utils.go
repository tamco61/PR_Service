package utils

import "app/internal/dto"

func JSONError(code dto.ErrorCode, message string) dto.ErrorResponese {
	return dto.ErrorResponese{
		Error: dto.ErrorDetail{
			Code:    code,
			Message: message,
		},
	}
}
