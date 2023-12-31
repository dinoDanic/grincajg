package models

import (
	"grincajg/database"
	"time"
)

type Organization struct {
	ID   uint   `gorm:"primarykey"`
	Name string `gorm:"type:varchar(100);not null"`

	AdminUserID uint `gorm:"unique;index"`
	AdminUser   User `gorm:"foreignKey:AdminUserID"`
	Users       []User
	Stores      []Store

	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateOrganizationInput struct {
  Name string `json:"name" validate:"required" example:"organization example 1"`
}

type OrganizationRecord struct {
	ID   uint   `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func FilterOrganizationRecord(organization Organization) OrganizationRecord {
	return OrganizationRecord{
		Name: organization.Name,
		ID:   organization.ID,
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
