package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mukhtar-husnain/go-gin/controller"
	"github.com/mukhtar-husnain/go-gin/service"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default()

	server.GET("/videos", func(c *gin.Context) {
		videos := videoController.FindAll()
		c.JSON(http.StatusOK, videos)
	})

	server.POST("/videos", func(c *gin.Context) {
		videos := videoController.Save(c)
		c.JSON(http.StatusOK, videos)
	})
	server.Run(":8080")
}
