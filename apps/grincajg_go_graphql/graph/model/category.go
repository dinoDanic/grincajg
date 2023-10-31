package model

import (
	"time"
)

type Category struct {
	ID         int    `gorm:"primarykey"`
	Name       string `gorm:"type:varchar(100);not null;unique"`
	CategoryID int
	Parent     *Category `gorm:"foreignKey:CategoryID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
