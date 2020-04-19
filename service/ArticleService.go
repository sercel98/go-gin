package service

import (
	"errors"
	"github.com/sercel98/go-gin/models"
)

type ArticleService interface {
	SaveArticle(article models.Article) models.Article
	FindAllArticles() []models.Article
	FindOneArticle(title string) (models.Article, error)
}

//Struct - Implements ArticleService Interface
type articleService struct {
	articles []models.Article
}

func NewArticleService() ArticleService {
	return &articleService{}
}

func (a *articleService) SaveArticle(article models.Article) models.Article {
	a.articles = append(a.articles, article)
	return article
}

func (a *articleService) FindAllArticles() []models.Article {
	return a.articles
}

func (a *articleService) FindOneArticle(title string) (models.Article, error) {

	var article models.Article
	for _, art := range a.articles {
		if art.Title == title {
			article = art
			return article, nil
		}
	}
	return article, errors.New("Article not found")
}
