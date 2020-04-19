package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sercel98/go-gin/models"
	"github.com/sercel98/go-gin/service"
)

//Interface VideoController
type VideoController interface {
	FindAll() []models.Video
	Save(context *gin.Context) models.Video
}

//Returns the memory location of a new element of type VideoController Interface, receives the corresponding service
func NewController(s service.VideoService) VideoController {
	return &videoController{
		service: s,
	}
}

//Struct - Implements VideoController Interface
type videoController struct {
	service service.VideoService
}

//This method returns the array of videos
func (c *videoController) FindAll() []models.Video {
	return c.service.FindAll()
}

//This method saves a ne video. Receives the context containing the Video information in JSON format.
func (c *videoController) Save(context *gin.Context) models.Video {
	var video models.Video
	context.BindJSON(&video)
	c.service.Save(video)
	return video
}
