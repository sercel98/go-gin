package service

import (
	"errors"
	"github.com/sercel98/go-gin/entity"
)

type ArticleService interface {
	SaveArticle(article entity.Article) entity.Article
	FindAllArticles() []entity.Article
	FindOneArticle(title string) (entity.Article, error)
}

//Struct - Implements ArticleService Interface
type articleService struct {
	articles []entity.Article
}

func NewArticleService() ArticleService {
	return &articleService{}
}

func (a *articleService) SaveArticle(article entity.Article) entity.Article {
	a.articles = append(a.articles, article)
	return article
}

func (a *articleService) FindAllArticles() []entity.Article {
	return a.articles
}

func (a *articleService) FindOneArticle(title string) (entity.Article, error) {

	var article entity.Article
	for _, art := range a.articles {
		if art.Title == title {
			article = art
			return article, nil
		}
	}
	return article, errors.New("Article not found")
}
