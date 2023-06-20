package service

import (
	"go.uber.org/zap"
	"hexagonal/application/domain"
	"hexagonal/configuration/logger"
	"hexagonal/configuration/rest_errors"
)

func (ud *userDomainService) FindUserByIDServices(
	id string,
) (*domain.UserDomain, *rest_errors.RestErr) {
	logger.Info("Init findUserByID services.",
		zap.String("journey", "findUserById"))

	return ud.repository.FindUserByID(id)
}

func (ud *userDomainService) FindUserByEmailServices(
	email string,
) (*domain.UserDomain, *rest_errors.RestErr) {
	logger.Info("Init findUserByEmail services.",
		zap.String("journey", "findUserById"))

	return ud.repository.FindUserByEmail(email)
}
