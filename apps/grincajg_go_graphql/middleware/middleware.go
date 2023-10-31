package middleware

import (
	"context"
	"fmt"

	"grincajg/database"
	"grincajg/graph/model"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/vektah/gqlparser/v2/gqlerror"
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
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)

			if ok == false {
				next.ServeHTTP(w, r)
				http.Error(w, "Invalid claims", http.StatusForbidden)
				return
			}

			var user model.User

			database.DB.First(&user, "id = ?", fmt.Sprint(claims["sub"]))

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

func GetUser(ctx context.Context) (*model.User, error) {
	user := ForContext(ctx)
	if user == nil {
		return &model.User{}, gqlerror.Errorf("Not authorized")
	}

	return user, nil
}

func GetSuperUser(ctx context.Context) (*model.User, error) {
	user := ForContext(ctx)
	if user == nil {
		return nil, gqlerror.Errorf("Not authorized")
	}
	if user.SuperAdmin == nil || !*user.SuperAdmin {
		return nil, gqlerror.Errorf("No permission")
	}
	return user, nil
}
