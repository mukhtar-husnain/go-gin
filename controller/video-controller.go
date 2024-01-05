package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mukhtar-husnain/go-gin/entity"
	"github.com/mukhtar-husnain/go-gin/service"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(c *gin.Context) error
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

func (controller *controller) FindAll() []entity.Video {
	return controller.service.FindAll()
}

func (controller *controller) Save(c *gin.Context) error {
	var video entity.Video
	err := c.ShouldBindJSON(&video)
	if err != nil {
		return err
	}
	controller.service.Save(video)
	return nil
}
