package main

import (
	"github.com/HunCoding/golang-architecture/mvc/src/controller"
	"github.com/HunCoding/golang-architecture/mvc/src/controller/repository"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	client, _ := mongo.NewClient()
	repo := repository.NewUserRepository(client)
	control := controller.NewUserController(repo)

	router := gin.Default()

	router.POST("/user", control.CreateUser)

	router.Run(":9090")
}
