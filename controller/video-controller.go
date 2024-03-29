package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/mukhtar-husnain/go-gin/entity"
	"github.com/mukhtar-husnain/go-gin/service"
	"github.com/mukhtar-husnain/go-gin/validators"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(c *gin.Context) error
	ShowAll(c *gin.Context)
}

type controller struct {
	service service.VideoService
}

var validate *validator.Validate

func New(service service.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
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
	err = validate.Struct(video)
	if err != nil {
		return err
	}
	controller.service.Save(video)
	return nil
}

func (controller *controller) ShowAll(c *gin.Context) {
	videos := controller.service.FindAll()

	data := gin.H{
		"title":  "Video Page",
		"videos": videos,
	}

	c.HTML(http.StatusOK, "index.html", data)
}
