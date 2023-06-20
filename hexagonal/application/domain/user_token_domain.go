package domain

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"hexagonal/application/constants"
	"hexagonal/configuration/rest_errors"
	"os"
	"time"
)

func (ud *UserDomain) GenerateToken() (string, *rest_errors.RestErr) {
	secret := os.Getenv(constants.JWT_SECRET_KEY)

	claims := jwt.MapClaims{
		"id":    ud.Id,
		"email": ud.Email,
		"name":  ud.Name,
		"age":   ud.Age,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", rest_errors.NewInternalServerError(
			fmt.Sprintf("error trying to generate jwt token, err=%s", err.Error()))
	}

	return tokenString, nil
}
