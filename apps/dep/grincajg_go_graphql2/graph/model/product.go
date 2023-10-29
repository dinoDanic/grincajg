package model

import "time"

type Product struct {
	ID         uint   `gorm:"primarykey"`
	Name       string `gorm:"type:varchar(100);not null;unique"`
	CategoryID uint
	Category   *Category `gorm:"foreignKey:CategoryID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
