package util

import "github.com/go-playground/validator/v10"

func ErrorValidation(request interface{}) error {
	validate := validator.New()
	return validate.Struct(request)
}
