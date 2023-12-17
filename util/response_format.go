package util

import "go-short-url/internal/model/dto"

func ResponseFormat(code int, msg string, data any) dto.APIResponse {
	response := dto.APIResponse{
		Code:    code,
		Message: msg,
		Data:    data,
	}
	return response
}
