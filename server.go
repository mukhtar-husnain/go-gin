package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mukhtar-husnain/go-gin/controller"
	"github.com/mukhtar-husnain/go-gin/middleware"
	"github.com/mukhtar-husnain/go-gin/service"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	setupLogOutput()
	server := gin.New()

	server.Use(gin.Recovery(), middleware.Logger(), 
	middleware.BasicAuth(), gindump.Dump())

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
