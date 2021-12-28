package controller

import (
	"net/http"
	"pampam/backend/tuqa/entity"
	"pampam/backend/tuqa/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	FindAll() []entity.User
	Save(ctx *gin.Context) error
	Show(ctx *gin.Context)
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
	Login(ctx *gin.Context) entity.User
}

type userController struct {
	service service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &userController{
		service: service,
	}
}

func (c *userController) FindAll() []entity.User {
	return c.service.FindAll()
}

func (c *userController) Save(ctx *gin.Context) error {
	var user entity.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		return err
	}
	c.service.Save(user)
	return nil
}
func (c *userController) Show(ctx *gin.Context) {
	device := c.service.FindAll()
	ctx.HTML(http.StatusOK, "index.html", device)
}
func (c *userController) Update(ctx *gin.Context) error {
	var user entity.User
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	user.Id = id
	c.service.Delete(user)
	return nil
}
func (c *userController) Delete(ctx *gin.Context) error {
	var user entity.User
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	user.Id = id
	c.service.Delete(user)
	return nil
}
func (c *userController) Login(ctx *gin.Context) entity.User {
	var user entity.User
	ctx.ShouldBindJSON(&user)
	return c.service.Login(user)
}
