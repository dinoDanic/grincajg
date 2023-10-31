package model

import (
	"grincajg/database"
	"time"
)

type Store struct {
	ID             int    `gorm:"primarykey"`
	Name           string `gorm:"type:varchar(100);not null;unique"`
	Address        string `gorm:"type:varchar(100);not null"`
	OrganizationID int    `gorm:"foreignKey:Organization"`
	Organization   Organization

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (store *Store) SaveStore() (*Store, error) {
	err := database.DB.Create(&store).Error
	if err != nil {
		return &Store{}, err
	}
	return store, nil
}
