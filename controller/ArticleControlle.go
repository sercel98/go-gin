package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sercel98/go-gin/entity"
	service2 "github.com/sercel98/go-gin/service"
)

type ArticleController interface {
	Save(context *gin.Context) entity.Article
	FindAll() []entity.Article
}

type articleController struct {
	service service2.ArticleService
}

func NewArticleController(s service2.ArticleService) ArticleController {
	return &articleController{service: s}
}

func (a *articleController) Save(ctx *gin.Context) entity.Article {
	var article entity.Article
	ctx.BindJSON(&article)
	a.service.SaveArticle(article)
	return article
}

func (a *articleController) FindAll() []entity.Article {
	return a.service.FindAllArticles()
}
