package models

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name       string `gorm:"type:varchar(100);not null"`
	Email      string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password   string `gorm:"type:varchar(100);not null"`
	Verified   *bool  `gorm:"not null;default:false"`
	SuperAdmin *bool  `gorm:"not null;default:false"`

	OrganizationID *uint `gorm:"index"`
	Organization   *Organization
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
	}
}

func GetContextUser(c *fiber.Ctx) *User {
	user := c.Locals("user").(*User)
	return user
}
