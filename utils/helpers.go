package utils

import (
	"log"
	"os"

	"goauthapi/models"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.Fatal(err)
	}

	return string(hash)
}

func ComparePasswords(hashedPassword, plainPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	if err != nil {
		return err
	}

	return nil
}

func GenerateToken(user models.User) (string, error) {
	secret := os.Getenv("JWT_SECTET")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"iss": "protected",
	})

	genToken, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Fatal(err)
	}

	return genToken, nil
}
