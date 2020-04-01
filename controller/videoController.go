package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sercel98/go-gin/entity"
	"github.com/sercel98/go-gin/service"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(context *gin.Context) entity.Video
}

type controller struct {
	service service.VideoService
}

func NewController(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *controller) Save(context *gin.Context) entity.Video {
	var video entity.Video
	context.BindJSON(&video)
	c.service.Save(video)
	return video
}
