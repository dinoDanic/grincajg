package model

import (
	"grincajg/database"
	"time"

	"github.com/google/uuid"
)

type Organization struct {
	ID        *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name      string     `gorm:"type:varchar(100);not null"`
	UserID    *uuid.UUID
	CreatedAt *time.Time `gorm:"not null;default:now()"`
	UpdatedAt *time.Time `gorm:"not null;default:now()"`
}

type CreateOrganizationInput struct {
	Name string `json:"name"  validate:"required"`
}

func (organization *Organization) Save() (*Organization, error) {
	err := database.DB.Create(&organization).Error
	if err != nil {
		return &Organization{}, err
	}
	return organization, nil
}
