package views

import "github.com/HunCoding/golang-architecture/mvc/src/models"

type User struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
}

func ConvertUserIntoView(user *models.User) *User {
	return &User{
		Email: user.Email,
		Name:  user.Name,
		Age:   user.Age,
	}
}
