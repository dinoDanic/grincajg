package models

import (
	"gorm.io/gorm"
	"grincajg/database"
)

type Store struct {
	gorm.Model
	Name           string `gorm:"type:varchar(100);not null;unique"`
	Address        string `gorm:"type:varchar(100);not null"`
	OrganizationID uint   `gorm:"foreignKey:Organization"`
	Organization   Organization
}

type CreateStoreInput struct {
	Name    string `json:"name"  validate:"required"`
	Address string `json:"address"  validate:"required"`
}

type StoreRecord struct {
	ID      uint   `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Address string `json:"address,omitempty"`
}

func FilterStoreRecord(store Store) StoreRecord {
	return StoreRecord{
		Name:    store.Name,
		Address: store.Address,
		ID:      store.ID,
	}
}

func FilterStoreRecords(stores []Store) []StoreRecord {
	var storeRecords []StoreRecord
	for _, store := range stores {
		record := FilterStoreRecord(store)
		storeRecords = append(storeRecords, record)
	}
	return storeRecords
}

func (store *Store) SaveStore() (*Store, error) {
	err := database.DB.Create(&store).Error
	if err != nil {
		return &Store{}, err
	}
	return store, nil
}
