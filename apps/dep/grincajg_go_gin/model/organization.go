package model

import (
	database "grincajg_api/databse"

	"gorm.io/gorm"
)

type Organization struct {
	gorm.Model
	Name    string `gorm:"size:255;not null;unique" json:"name"`
	Address string `gorm:"size:255;not null;unique" json:"address"`
	UserID  uint
}

func (organization *Organization) Save() (*Organization, error) {
	err := database.Database.Create(&organization).Error
	if err != nil {
		return &Organization{}, err
	}
	return organization, nil
}
