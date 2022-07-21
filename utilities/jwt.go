package utilities

import (
	"aims/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const SecretKey = "simple_auth"

func CreateJWTToken(user models.User) (string, int64, error) {
	exp := time.Now().Add(time.Minute * 30).Unix()
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.Id
	claims["user_email"] = user.Email
	claims["user_username"] = user.Username
	claims["exp"] = exp
	t, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", 0, err
	}

	return t, exp, nil
}
