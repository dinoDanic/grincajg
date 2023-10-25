package models

import (
	"gorm.io/gorm"
	"grincajg/database"
)

type Organization struct {
	gorm.Model
	Name   string `gorm:"type:varchar(100);not null"`
	UserID uint   `gorm:"unique;foreignKey:User"`
	User   User
}

type CreateOrganizationInput struct {
	Name string `json:"name"  validate:"required"`
}

type OrganizationResponse struct {
	Name string `json:"name,omitempty"`
}

func FilterOrganizationResponse(organization Organization) OrganizationResponse {
	return OrganizationResponse{
		Name: organization.Name,
	}
}

func (organization *Organization) SaveOrganization() (*OrganizationResponse, error) {
	err := database.DB.Create(&organization).Error
	if err != nil {
		return &OrganizationResponse{}, err
	}

	filterResponse := FilterOrganizationResponse(*organization)

	return &filterResponse, nil

}
