package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"hexagonal/application/constants"
	"hexagonal/application/domain"
	"hexagonal/configuration/logger"
	"hexagonal/configuration/rest_errors"
	"os"
	"strings"
)

func VerifyTokenMiddleware(c *gin.Context) {
	secret := os.Getenv(constants.JWT_SECRET_KEY)
	tokenValue := RemoveBearerPrefix(c.Request.Header.Get("Authorization"))

	token, err := jwt.Parse(tokenValue, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}

		return nil, rest_errors.NewBadRequestError("invalid token")
	})
	if err != nil {
		errRest := rest_errors.NewUnauthorizedRequestError("invalid token")
		c.JSON(errRest.Code, errRest)
		c.Abort()
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		errRest := rest_errors.NewUnauthorizedRequestError("invalid token")
		c.JSON(errRest.Code, errRest)
		c.Abort()
		return
	}

	userDomain := domain.UserDomain{
		Id:    claims["id"].(string),
		Email: claims["email"].(string),
		Name:  claims["name"].(string),
		Age:   int8(claims["age"].(float64)),
	}
	logger.Info(fmt.Sprintf("User authenticated: %#v", userDomain))
}

func RemoveBearerPrefix(token string) string {
	if strings.HasPrefix(token, "Bearer ") {
		token = strings.TrimPrefix("Bearer ", token)
	}

	return token
}
