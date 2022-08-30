package repository

import (
	"github.com/HunCoding/golang-architecture/mvc/src/models"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRepository(
	mongodbClient *mongo.Client,
) UserRepository {
	return &userRepository{mongodbClient}
}

type UserRepository interface {
	CreateUser(user *models.User) (*models.User, error)
}

type userRepository struct {
	mongodbClient *mongo.Client
}

func (ur *userRepository) CreateUser(user *models.User) (*models.User, error) {
	//TODO: Implement mongodb user insert
	return nil, nil
}
