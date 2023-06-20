package service

import (
	"go.uber.org/zap"
	"hexagonal/application/domain"
	"hexagonal/application/port/input"
	"hexagonal/application/port/output"
	"hexagonal/configuration/logger"
	"hexagonal/configuration/rest_errors"
)

func NewUserDomainService(
	userRepository output.UserPort) input.UserDomainService {
	return &userDomainService{
		userRepository,
	}
}

type userDomainService struct {
	repository output.UserPort
}

func (ud *userDomainService) CreateUserServices(
	userDomain domain.UserDomain,
) (*domain.UserDomain, *rest_errors.RestErr) {

	logger.Info("Init createUser model.",
		zap.String("journey", "createUser"))

	user, _ := ud.FindUserByEmailServices(userDomain.Email)
	if user != nil {
		return nil, rest_errors.NewBadRequestError("Email is already registered in another account")
	}

	userDomain.EncryptPassword()
	userDomainRepository, err := ud.repository.CreateUser(userDomain)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "createUser"))
		return nil, err
	}

	logger.Info(
		"CreateUser service executed successfully",
		zap.String("userId", userDomainRepository.Id),
		zap.String("journey", "createUser"))
	return userDomainRepository, nil
}
