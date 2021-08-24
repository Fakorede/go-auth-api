package middlewares

import (
	"errors"
	"fmt"
	"goauthapi/utils"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func VerifyTokenMiddleware(next http.HandlerFunc) http.HandlerFunc {
	secret := os.Getenv("JWT_SECTET")

	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")

		if len(bearerToken) == 2 {
			authToken := bearerToken[1]

			token, error := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("an error occured")
				}
				return []byte(secret), nil
			})

			if error != nil {
				utils.ErrorJSON(rw, http.StatusUnauthorized, error)
				return
			}

			if token.Valid {
				next.ServeHTTP(rw, r)
			} else {
				utils.ErrorJSON(rw, http.StatusUnauthorized, error)
				return
			}

		} else {
			utils.ErrorJSON(rw, http.StatusUnauthorized, errors.New("invalid token"))
			return
		}
	})
}
