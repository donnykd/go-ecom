package auth

import (
	"strconv"
	"time"

	"github.com/donnykd/go-ecom/config"
	"github.com/golang-jwt/jwt"
)

func CreateJWT(secret []byte, userID int) (string, error) {
	expiration := time.Second * time.Duration(config.InitConfig().JWTExpirationInSeconds)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":    strconv.Itoa(userID),
		"expired_at": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
