package model

import "time"

type Product struct {
	ID         int   `gorm:"primarykey"`
	Name       string `gorm:"type:varchar(100);not null;unique"`
	CategoryID int
	Category   *Category `gorm:"foreignKey:CategoryID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
