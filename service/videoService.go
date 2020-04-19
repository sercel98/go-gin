package service

import "github.com/sercel98/go-gin/models"

//Interface
type VideoService interface {
	Save(models.Video) models.Video
	FindAll() []models.Video
}

//Struct - Implements VideoService Interface
type videoService struct {
	videos []models.Video
}

//Returns the memory location of a new element of type VideoService.
func NewVideoService() VideoService {
	return &videoService{}
}

//Adds an Video to the slice of videos
func (service *videoService) Save(video models.Video) models.Video {
	service.videos = append(service.videos, video)
	return video
}

//Returns the array of videos
func (service *videoService) FindAll() []models.Video {
	return service.videos
}
