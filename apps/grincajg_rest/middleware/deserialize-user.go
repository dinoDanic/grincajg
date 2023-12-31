package middleware

import (
	"fmt"
	"grincajg/database"
	"grincajg/models"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func DeserializeUser(c *fiber.Ctx) error {
	var tokenString string
	authorization := c.Get("Authorization")

	if strings.HasPrefix(authorization, "Bearer ") {
		tokenString = strings.TrimPrefix(authorization, "Bearer ")
	} else if c.Cookies("token") != "" {
		tokenString = c.Cookies("token")
	}

	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "You are not logged in"})
	}

	tokenByte, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
		}

		JWT_SECRET := os.Getenv("JWT_SECRET")

		return []byte(JWT_SECRET), nil
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": fmt.Sprintf("invalidate token: %v", err)})
	}

	claims, ok := tokenByte.Claims.(jwt.MapClaims)
	if !ok || !tokenByte.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "invalid token claim"})

	}

	var user models.User
	database.DB.First(&user, "id = ?", fmt.Sprint(claims["sub"]))

	userIdStr := fmt.Sprintf("%d", user.ID)
	claimsSubStr := fmt.Sprintf("%.0f", claims["sub"].(float64))

	if userIdStr != claimsSubStr {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": "The user belonging to this token no longer exists"})
	}

	c.Locals("user", &user)

	return c.Next()
}
