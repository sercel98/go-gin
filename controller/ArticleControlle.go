package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sercel98/go-gin/models"
	service2 "github.com/sercel98/go-gin/service"
)

type ArticleController interface {
	Save(context *gin.Context) models.Article
	FindAll() []models.Article
}

type articleController struct {
	service service2.ArticleService
}

func NewArticleController(s service2.ArticleService) ArticleController {
	return &articleController{service: s}
}

func (a *articleController) Save(ctx *gin.Context) models.Article {
	var article models.Article
	ctx.BindJSON(&article)
	a.service.SaveArticle(article)
	return article
}

func (a *articleController) FindAll() []models.Article {
	return a.service.FindAllArticles()
}
