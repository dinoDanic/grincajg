package model

import (
	"time"
)

type User struct {
	ID         string `gorm:"primarykey"`
	Name       string `gorm:"type:varchar(100);not null"`
	Email      string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password   string `gorm:"type:varchar(100);not null"`
	Verified   *bool  `gorm:"not null;default:false"`
	SuperAdmin *bool  `gorm:"not null;default:false"`

	OrganizationID *int `gorm:"index"`
	Organization   *Organization

	CreatedAt time.Time
	UpdatedAt time.Time
}

type SignUpInput struct {
	Name            string `json:"name" validate:"required" example:"account1"`
	Email           string `json:"email" validate:"required,email" example:"account1@mail.com"`
	Password        string `json:"password" validate:"required" example:"Ruda,actv1!"`
	PasswordConfirm string `json:"passwordConfirm" validate:"required" example:"Ruda,actv1!"`
}

type SignInInput struct {
	Email    string `json:"email" validate:"required" example:"account1@mail.com"`
	Password string `json:"password" validate:"required" example:"Ruda,actv1!"`
}

type SignInResponse struct {
	Status string `json:"status" example:"success"`
	Token  string `json:"token,omitempty" example:"eyJhbGciOiJIUzI1N..."`
}

type UserResponse struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Role  string `json:"role,omitempty"`
}

func FilterUserRecord(user *User) UserResponse {
	return UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}
