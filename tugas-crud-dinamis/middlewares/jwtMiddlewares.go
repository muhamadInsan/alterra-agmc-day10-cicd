package middlewares

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(userId int) (string, error) {

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() // token exp 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
