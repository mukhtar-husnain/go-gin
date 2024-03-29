package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mukhtar-husnain/go-gin/controller"
	"github.com/mukhtar-husnain/go-gin/middleware"
	"github.com/mukhtar-husnain/go-gin/service"
	// gindump "github.com/tpkeeper/gin-dump"
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

	server.Static("/css", "./templates/css")
	server.LoadHTMLGlob("templates/*.html")

	server.Use(gin.Recovery(), middleware.Logger(),
		middleware.BasicAuth())

	apiRoutes := server.Group("/api", middleware.BasicAuth())
	{
		apiRoutes.GET("/videos", func(c *gin.Context) {
			videos := videoController.FindAll()
			c.JSON(http.StatusOK, videos)
		})

		apiRoutes.POST("/videos", func(c *gin.Context) {
			err := videoController.Save(c)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, gin.H{"message": "valid"})
			}
		})
	}

	viewRoutes := server.Group("/view")
	viewRoutes.GET("/videos", videoController.ShowAll)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	server.Run(":" + port)
}
