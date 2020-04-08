package service

import (
	"errors"
	"github.com/sercel98/go-gin/entity"
)

type ArticleService interface {
	Save(article entity.Article) entity.Article
	FindAll() []entity.Article
	FindOne(title string) (entity.Article, error)
}

//Struct - Implements ArticleService Interface
type articleService struct {
	articles []entity.Article
}

func newArticleService() ArticleService  {
	return &articleService{}
}

func (a *articleService) Save(article entity.Article) entity.Article{
	a.articles = append(a.articles, article)
	return article
}

func (a *articleService) FindAll() []entity.Article{
	return a.articles
}

func (a *articleService) FindOne(title string) (entity.Article, error){

	var article entity.Article
	for _, art := range a.articles {
		if art.Title == title {
			article = art
			return article, nil
		}
	}
	return article,errors.New("Article not found")
}











//Returns the memory location of a new element of type VideoService.
func NewArticleService() ArticleService {
	return &videoService{}
}


//Adds an Video to the slice of videos
func (service *videoService) Save(video entity.Video) entity.Video {
	service.videos = append(service.videos, video)
	return video
}

//Returns the array of videos
func (service *videoService) FindAll() []entity.Video {
	return service.videos
}
