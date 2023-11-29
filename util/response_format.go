package util

import "go-short-url/internal/model/api"

func ResponseFormat(code int, msg string, data any) api.APIResponse {
	response := api.APIResponse{
		Code:    code,
		Message: msg,
		Data:    data,
	}
	return response
}
