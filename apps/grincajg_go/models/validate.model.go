package models

import "github.com/go-playground/validator/v10"

var validate = validator.New()

func ValidateStruct[T any](payload T) []*ValidateError {
	var errors []*ValidateError
	err := validate.Struct(payload)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ValidateError
			element.Field = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
