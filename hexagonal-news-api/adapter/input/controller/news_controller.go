package controller

import (
	"net/http"

	"github.com/HunCoding/golang-architecture/hexagonal-news-api/adapter/input/model/request"
	"github.com/HunCoding/golang-architecture/hexagonal-news-api/application/domain"
	"github.com/HunCoding/golang-architecture/hexagonal-news-api/application/port/input"
	"github.com/HunCoding/golang-architecture/hexagonal-news-api/configuration/logger"
	"github.com/HunCoding/golang-architecture/hexagonal-news-api/configuration/validation"
	"github.com/gin-gonic/gin"
)

type newsController struct {
	newsUseCase input.NewsUseCase
}

func NewNewsController(
	newsUseCase input.NewsUseCase,
) *newsController {
	return &newsController{newsUseCase}
}

func (nc *newsController) GetNews(c *gin.Context) {
	//q=tesla&from=2023-08-20&apiKey=df969581143d4958a630cb3d0efeaf21
	logger.Info("Init GetNews controller api")
	newsRequest := request.NewsRequest{}

	if err := c.ShouldBindQuery(&newsRequest); err != nil {
		logger.Error("Error trying to validate fields from request", err)
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	newsDomain := domain.NewsReqDomain{
		Subject: newsRequest.Subject,
		From:    newsRequest.From.Format("2006-01-02"),
	}

	newsResponseDomain, err := nc.newsUseCase.GetNewsService(newsDomain)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, newsResponseDomain)
}
