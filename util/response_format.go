package util

import "go-shorten/internal/model/dto"

func ResponseFormat(status int, msg string, data any) dto.APIResponse {
	response := dto.APIResponse{
		Status:  status,
		Message: msg,
		Data:    data,
	}
	return response
}
