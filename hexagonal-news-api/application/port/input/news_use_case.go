package input

import (
	"github.com/HunCoding/golang-architecture/hexagonal-news-api/application/domain"
	"github.com/HunCoding/golang-architecture/hexagonal-news-api/configuration/rest_err"
)

type NewsUseCase interface {
	GetNewsService(
		domain.NewsReqDomain,
	) (*domain.NewsDomain, *rest_err.RestErr)
}
