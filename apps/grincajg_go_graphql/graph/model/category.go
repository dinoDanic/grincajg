package model

import (
	"grincajg/database"
	"time"
)

type Category struct {
	ID         int    `gorm:"primarykey"`
	Name       string `gorm:"type:varchar(100);not null;unique"`
	CategoryID *int
	Parent     *Category

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (category *Category) Save() (*Category, error) {
	err := database.DB.Create(&category).Error
	if err != nil {
		return &Category{}, err
	}
	return category, nil
}
