package controller

import (
	"fmt"
	"net/http"

	"github.com/HunCoding/golang-architecture/mvc/src/models/repository"
	"github.com/HunCoding/golang-architecture/mvc/src/views"
	"github.com/gin-gonic/gin"
)

func NewUserController(
	userRepository repository.UserRepository,
) UserControllerInterface {
	return &userController{userRepository}
}

type UserControllerInterface interface {
	CreateUser(c *gin.Context)
}

type userController struct {
	userRepository repository.UserRepository
}

func (uc *userController) CreateUser(c *gin.Context) {
	user := UserForm{}

	if err := c.ShouldBindJSON(&user); err != nil {
		errPerso := views.ParseError(
			fmt.Sprintf("Error trying to parse user, error=%s", err),
			http.StatusBadRequest,
		)

		c.JSON(errPerso.ErrorCode, errPerso)
	}

	userModel := user.ConvertInputUserToUserEntity()
	createdUser, err := uc.userRepository.CreateUser(userModel)
	if err != nil {
		errPerso := views.ParseError(
			fmt.Sprintf("Error trying to create user, error=%s", err),
			http.StatusInternalServerError,
		)

		c.JSON(errPerso.ErrorCode, errPerso)
	}

	c.JSON(http.StatusOK, views.ConvertUserIntoView(createdUser))
}
