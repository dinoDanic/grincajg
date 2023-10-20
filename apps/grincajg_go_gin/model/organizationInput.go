package model

type OrganizationInput struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"password"`
}
