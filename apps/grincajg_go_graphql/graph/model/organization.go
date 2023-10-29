package model

import (
	"grincajg/database"
	"time"
)

type Organization struct {
	ID   int    `gorm:"primarykey"`
	Name string `gorm:"type:varchar(100);not null"`

	AdminUserID int  `gorm:"unique;index"`
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
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func (organization *Organization) SaveOrganization() (*Organization, error) {
	err := database.DB.Create(&organization).Error

	if err != nil {
		return &Organization{}, err
	}

	return organization, nil

}
