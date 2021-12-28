package controller

import (
	"pampam/backend/tuqa/entity"
	"pampam/backend/tuqa/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PointerController interface {
	FindMerge(ctx *gin.Context) []entity.Pointer
}

type pointerController struct {
	service service.PointerService
}

func NewPointerController(service service.PointerService) PointerController {
	return &pointerController{
		service: service,
	}
}

func (c *pointerController) FindMerge(ctx *gin.Context) []entity.Pointer {
	var user entity.User
	var err error
	user.Id, err = strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err == nil {
		return c.service.FindMerge(user)
	} else {
		return c.service.FindMerge(user)
	}
}
