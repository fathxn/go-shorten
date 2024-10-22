package util

import "go-shorten/internal/domain"

func ResponseFormat(status int, msg string, data any) domain.APIResponse {
	response := domain.APIResponse{
		Status:  status,
		Message: msg,
		Data:    data,
	}
	return response
}
