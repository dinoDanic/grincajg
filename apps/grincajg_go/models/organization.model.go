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
	Stores  []Store
}

type CreateOrganizationInput struct {
	Name string `json:"name"  validate:"required"`
}

type OrganizationRecord struct {
	Name string `json:"name,omitempty"`
}

func FilterOrganizationRecord(organization Organization) OrganizationRecord {
	return OrganizationRecord{
		Name: organization.Name,
	}
}

func (organization *Organization) SaveOrganization() (*OrganizationRecord, error) {
	err := database.DB.Create(&organization).Error
	if err != nil {
		return &OrganizationRecord{}, err
	}

	filterResponse := FilterOrganizationRecord(*organization)

	return &filterResponse, nil

}
