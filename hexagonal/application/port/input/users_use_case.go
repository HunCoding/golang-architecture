package input

import (
	"hexagonal/application/domain"
	"hexagonal/configuration/rest_errors"
)

type UserDomainService interface {
	CreateUserServices(domain.UserDomain) (
		*domain.UserDomain, *rest_errors.RestErr)
	FindUserByIDServices(
		id string,
	) (*domain.UserDomain, *rest_errors.RestErr)
	FindUserByEmailServices(
		email string,
	) (*domain.UserDomain, *rest_errors.RestErr)
}
