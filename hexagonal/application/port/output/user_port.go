package output

import (
	"hexagonal/application/domain"
	"hexagonal/configuration/rest_errors"
)

type UserPort interface {
	CreateUser(
		userDomain domain.UserDomain,
	) (*domain.UserDomain, *rest_errors.RestErr)

	FindUserByEmail(
		email string,
	) (*domain.UserDomain, *rest_errors.RestErr)
	FindUserByID(
		id string,
	) (*domain.UserDomain, *rest_errors.RestErr)
}
