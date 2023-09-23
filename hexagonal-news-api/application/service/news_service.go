package service

import (
	"fmt"

	"github.com/HunCoding/golang-architecture/hexagonal-news-api/application/domain"
	"github.com/HunCoding/golang-architecture/hexagonal-news-api/application/port/output"
	"github.com/HunCoding/golang-architecture/hexagonal-news-api/configuration/logger"
	"github.com/HunCoding/golang-architecture/hexagonal-news-api/configuration/rest_err"
)

type newsService struct {
	newsPort output.NewsPort
}

func NewNewsService(newsPort output.NewsPort) *newsService {
	return &newsService{newsPort}
}

func (ns *newsService) GetNewsService(
	newsDomain domain.NewsReqDomain,
) (*domain.NewsDomain, *rest_err.RestErr) {
	logger.Info(
		fmt.Sprintf("Init getNewsService function, subject=%s, from=%s",
			newsDomain.Subject, newsDomain.From))

	newsDomainResponse, err := ns.newsPort.GetNewsPort(newsDomain)
	return newsDomainResponse, err
}
