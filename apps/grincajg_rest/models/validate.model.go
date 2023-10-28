package models

import "github.com/go-playground/validator/v10"

var validate = validator.New()

type Error struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func ValidateStruct[T any](payload T) error {
	err := validate.Struct(payload)
	return err
}
