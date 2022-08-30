package controller

import "github.com/HunCoding/golang-architecture/mvc/src/models"

type UserForm struct {
	Email    string `json:"email" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Age      int    `json:"age" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (uf *UserForm) ConvertInputUserToUserEntity() *models.User {
	return &models.User{
		Email:    uf.Email,
		Name:     uf.Name,
		Age:      uf.Age,
		Password: uf.Password,
	}
}
