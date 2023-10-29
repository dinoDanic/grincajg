package jwt

import (
	"grincajg/database"
	"grincajg/graph/model"
	"log"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func ParseToken(tokenStr string) (string, error) {
	JWT_SECRET := os.Getenv("JWT_SECRET")

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWT_SECRET), nil
	})

	log.Println("kita")
	log.Println(tokenStr)

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		log.Println("tu sam")
		userId := claims["sub"].(string)
		log.Println("userId")
		log.Println(userId)
		return userId, nil
	} else {
		return "", err
	}
}

func GetUserByUserId(id string) (model.User, error) {

	var user model.User

	result := database.DB.First(&user, "id = ?", id)

	if result.Error != nil {
		return model.User{}, gqlerror.Errorf("Cant find user")
	}

	return user, nil

}
