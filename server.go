package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()

	server.GET("/test", func(cont *gin.Context))
}
