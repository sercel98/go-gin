package service

import "github.com/sercel98/go-gin/entity"

//Interface
type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}
//Returns the memory location of a new element of type VideoService.
func NewVideoService() VideoService {
	return &videoService{}
}

//Struct - Implements VideoService Interface
type videoService struct {
	videos []entity.Video
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

