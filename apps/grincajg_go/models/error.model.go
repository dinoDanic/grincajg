package models

type Error struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ValidateError struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value,omitempty"`
}

type ValidateErrorResponse struct {
	Status string `json:"status"`
	Errors []ValidateError
}
