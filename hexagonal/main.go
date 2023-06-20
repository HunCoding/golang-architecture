package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"hexagonal/adapter/input/controller"
	"hexagonal/adapter/input/controller/routes"
	"hexagonal/adapter/output/repository"
	service "hexagonal/application/services"
	"hexagonal/configuration/database/mongodb"
	"hexagonal/configuration/logger"
	"log"
)

func main() {
	logger.Info("About to start user application")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf(
			"Error trying to connect to database, error=%s \n",
			err.Error())
		return
	}

	userController := initDependencies(database)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func initDependencies(
	database *mongo.Database,
) controller.UserControllerInterface {
	userRepo := repository.NewUserRepository(database)
	userService := service.NewUserDomainService(userRepo)
	return controller.NewUserControllerInterface(userService)
}
