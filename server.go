package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sercel98/go-gin/controller"
	"github.com/sercel98/go-gin/service"
)

var (
	videoService    service.VideoService       = service.NewVideoService()
	videoController controller.VideoController = controller.NewController(videoService)
)

func main() {
	server := gin.Default()

	server.GET("/videos", func(context *gin.Context) {
		context.JSON(200, videoController.FindAll())
	})

	server.POST("/create", func(context *gin.Context) {
		context.JSON(200, videoController.Save(context))
	})
	server.Run(":5000")
}
