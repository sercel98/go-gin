package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sercel98/go-gin/entity"
	"github.com/sercel98/go-gin/service"
)

//Interface VideoController
type VideoController interface {
	FindAll() []entity.Video
	Save(context *gin.Context) entity.Video
}

//Returns the memory location of a new element of type VideoController Interface, receives the corresponding service
func NewController(service service.VideoService) VideoController {
	return &videoController{
		service: service,
	}
}

//Struct - Implements VideoController Interface
type videoController struct {
	service service.VideoService
}

//This method returns the array of videos
func (c *videoController) FindAll() []entity.Video {
	return c.service.FindAll()
}

//This method saves a ne video. Receives the context containing the Video information in JSON format.
func (c *videoController) Save(context *gin.Context) entity.Video {
	var video entity.Video
	context.BindJSON(&video)
	c.service.Save(video)
	return video
}
