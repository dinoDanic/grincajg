package models

import "github.com/go-playground/validator/v10"

var validate = validator.New()

type ValidateErrorResponse struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value,omitempty"`
}

type ErrorResponse struct {
  Status  string `json:"status"`
  Message string `json:"message"`
}

func ValidateStruct[T any](payload T) []*ValidateErrorResponse {
	var errors []*ValidateErrorResponse
	err := validate.Struct(payload)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ValidateErrorResponse
			element.Field = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
