package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	c "github.com/sercel98/go-gin/config"
	"github.com/sercel98/go-gin/controller"
	"github.com/sercel98/go-gin/database"
	"github.com/sercel98/go-gin/service"
	"strconv"
)

//Global variables for controllers
var (
	videoService      service.VideoService       = service.NewVideoService()
	videoController   controller.VideoController = controller.NewController(videoService)
	articleService                               = service.NewArticleService()
	articleController                            = controller.NewArticleController(articleService)
)

func main() {




	config := c.NewConfiguration()

	db := database.NewPgConnection(config)
	//Define a close operation at the end of the method
	db.CreateTables()

	server := gin.Default()
	server.GET("/videos", func(context *gin.Context) {
		context.JSON(200, videoController.FindAll())
	})

	server.POST("/create", func(context *gin.Context) {
		context.JSON(200, videoController.Save(context))
	})
	server.Run(":" + strconv.Itoa(config.Server.Port))
}
