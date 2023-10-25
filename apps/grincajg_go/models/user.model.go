package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string  `gorm:"type:varchar(100);not null"`
	Email    string  `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password string  `gorm:"type:varchar(100);not null"`
	Role     *string `gorm:"type:varchar(50);default:'user';not null"`
	Verified *bool   `gorm:"not null;default:false"`
}

type SignUpInput struct {
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required"`
	PasswordConfirm string `json:"passwordConfirm" validate:"required"`
}

type SignInInput struct {
	Email    string `json:"email"  validate:"required"`
	Password string `json:"password"  validate:"required"`
}

type UserResponse struct {
	ID    uint   `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Role  string `json:"role,omitempty"`
}

func FilterUserRecord(user *User) UserResponse {
	return UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  *user.Role,
	}
}
