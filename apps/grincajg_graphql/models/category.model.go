package models

import (
	"time"
)

type Category struct {
	ID         uint   `gorm:"primarykey"`
	Name       string `gorm:"type:varchar(100);not null;unique"`
	CategoryID uint
	Parent     *Category `gorm:"foreignKey:CategoryID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
