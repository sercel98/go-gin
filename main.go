package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	c "github.com/sercel98/go-gin/config"
	"github.com/sercel98/go-gin/controller"
	"github.com/sercel98/go-gin/service"
	"github.com/spf13/viper"
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

	//viper is used to read a config.yml file where is specified the server's port and the database configutarion
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	viper.SetConfigType("yml")

	//Create a variable for the interface Configurations, later it will be used to unmarshal the yml
	var config c.Configurations

	//Read the config file. Check for errors also.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	//decode the yml into the interface
	err := viper.Unmarshal(&config)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	//Connect to database
	//TODO: Make this a struct and implement singleton
	db, err := gorm.Open("postgres", "host="+config.Database.DBHost+" port=5432 user="+
		config.Database.DBUser+" dbname="+config.Database.DBName+" password="+config.Database.DBName)

	//Define a close operation at the end of the method
	defer db.Close()

	server := gin.Default()
	server.GET("/videos", func(context *gin.Context) {
		context.JSON(200, videoController.FindAll())
	})

	server.POST("/create", func(context *gin.Context) {
		context.JSON(200, videoController.Save(context))
	})
	server.Run(":" + strconv.Itoa(config.Server.Port))
}
