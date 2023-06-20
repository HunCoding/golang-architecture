package converter

import (
	"hexagonal/adapter/output/model/entity"
	"hexagonal/application/domain"
)

func ConvertDomainToEntity(
	domain domain.UserDomain,
) *entity.UserEntity {
	return &entity.UserEntity{
		Email:    domain.Email,
		Password: domain.Password,
		Name:     domain.Name,
		Age:      domain.Age,
	}
}
