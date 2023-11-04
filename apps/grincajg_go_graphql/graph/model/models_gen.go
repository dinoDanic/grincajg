// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CreateCategoryInput struct {
	Name             string  `json:"name"`
	ParentCategoryID *string `json:"parentCategoryId,omitempty"`
}

type CreateOrganizationInput struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type CreateStoreInput struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type CreateUserInput struct {
	Email           string `json:"email"`
	Name            string `json:"name"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
}

type LoginUserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Session struct {
	Token string `json:"token"`
}
