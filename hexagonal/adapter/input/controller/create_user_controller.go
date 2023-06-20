package controller

import (
	"hexagonal/adapter/input/converter"
	"hexagonal/adapter/input/model/request"
	"hexagonal/application/domain"
	"hexagonal/application/port/input"
	"hexagonal/configuration/logger"
	"hexagonal/configuration/validation"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func NewUserControllerInterface(
	serviceInterface input.UserDomainService,
) UserControllerInterface {
	return &userControllerInterface{
		service: serviceInterface,
	}
}

type UserControllerInterface interface {
	FindUserByID(c *gin.Context)
	FindUserByEmail(c *gin.Context)

	CreateUser(c *gin.Context)
}

type userControllerInterface struct {
	service input.UserDomainService
}

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zap.String("journey", "createUser"),
	)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "createUser"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	userDomain := domain.UserDomain{
		Email:    userRequest.Email,
		Password: userRequest.Password,
		Name:     userRequest.Name,
		Age:      userRequest.Age,
	}
	domainResult, err := uc.service.CreateUserServices(userDomain)
	if err != nil {
		logger.Error(
			"Error trying to call CreateUser service",
			err,
			zap.String("journey", "createUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"CreateUser controller executed successfully",
		zap.String("userId", domainResult.Id),
		zap.String("journey", "createUser"))

	c.JSON(http.StatusOK, converter.ConvertDomainToResponse(
		domainResult,
	))
}
