package controllers

import (
	"context"
	"grincajg/database"
	"grincajg/graph/model"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {

	if input.Pasword != input.PosswordConfirm {
		return &model.User{}, gqlerror.Errorf("Passwords do not match")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Pasword), bcrypt.DefaultCost)
	//
	if err != nil {
		return &model.User{}, gqlerror.Errorf("Problem with password")
	}
	//
	newUser := model.User{
		Name:     input.Name,
		Email:    strings.ToLower(input.Email),
		Password: string(hashedPassword),
	}

	result := database.DB.Create(&newUser)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		return &model.User{}, gqlerror.Errorf("User with that email already exists")
	} else if result.Error != nil {
		return &model.User{}, gqlerror.Errorf("Something went wrong")
	}

	return &newUser, nil
}

func LoginUser(ctx context.Context, input model.LoginUserInput) (*model.Session, error) {

	var user model.User

	result := database.DB.First(&user, "email = ?", strings.ToLower(input.Email))

	if result.Error != nil {
		return &model.Session{}, gqlerror.Errorf("Invalid email or Password")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))

	if err != nil {
		return &model.Session{}, gqlerror.Errorf("Invalid email or Password")
	}

	JWT_SECRET := os.Getenv("JWT_SECRET")

	tokenByte := jwt.New(jwt.SigningMethodHS256)

	now := time.Now().UTC()

	claims := tokenByte.Claims.(jwt.MapClaims)

	claims["sub"] = user.ID
	claims["exp"] = now.Add(time.Hour).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	tokenString, err := tokenByte.SignedString([]byte(JWT_SECRET))

	if err != nil {
		return &model.Session{}, gqlerror.Errorf("generating JWT Token failed")
	}

	return &model.Session{Token: tokenString}, nil
}
