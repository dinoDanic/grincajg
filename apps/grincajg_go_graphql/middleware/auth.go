package auth

import (
	"context"
	"fmt"

	// "grincajg/database"
	"grincajg/database"
	"grincajg/graph/model"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func Middleware() func(http.Handler) http.Handler {

	JWT_SECRET := os.Getenv("JWT_SECRET")
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")

			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			tokenStr := header

			token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(JWT_SECRET), nil
			})

			if err != nil {
				fmt.Println("Error parsing token:", err.Error())
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)

			if ok == false {
				next.ServeHTTP(w, r)
				return
			}

			var user model.User

			database.DB.First(&user, "id = ?", fmt.Sprint(claims["sub"]))

			userIdStr := fmt.Sprintf("%d", user.ID)
			claimsSubStr := fmt.Sprintf("%.0f", claims["sub"].(float64))

			if userIdStr != claimsSubStr {
				fmt.Println("Error parsing token:", "The user belonging to this token no longer exists")
			}

			ctx := context.WithValue(r.Context(), userCtxKey, &user)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)

		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *model.User {
	raw, _ := ctx.Value(userCtxKey).(*model.User)
	return raw
}
