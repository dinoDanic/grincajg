package models

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID         uint   `gorm:"primarykey"`
	Name       string `gorm:"type:varchar(100);not null"`
	Email      string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password   string `gorm:"type:varchar(100);not null"`
	Verified   *bool  `gorm:"not null;default:false"`
	SuperAdmin *bool  `gorm:"not null;default:false"`

	OrganizationID *uint `gorm:"index"`
	Organization   *Organization

	CreatedAt time.Time
	UpdatedAt time.Time
}

type SignUpInput struct {
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required"`
	PasswordConfirm string `json:"passwordConfirm" validate:"required"`
}

type SignInInput struct {
	Email    string `json:"email" validate:"required" example:"account1@mail.com"`
	Password string `json:"password" validate:"required" example:"1"`
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
	}
}

func GetContextUser(c *fiber.Ctx) *User {
	user := c.Locals("user").(*User)
	return user
}
