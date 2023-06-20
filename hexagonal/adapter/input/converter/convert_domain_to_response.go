package converter

import (
	"hexagonal/adapter/input/model/response"
	"hexagonal/application/domain"
)

func ConvertDomainToResponse(
	userDomain *domain.UserDomain,
) response.UserResponse {
	return response.UserResponse{
		ID:    userDomain.Id,
		Email: userDomain.Email,
		Name:  userDomain.Name,
		Age:   userDomain.Age,
	}
}
