package controller

import (
	"net/http"
	"pampam/backend/tuqa/entity"
	"pampam/backend/tuqa/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeviceController interface {
	FindAll() []entity.Device
	Save(ctx *gin.Context) error
	Show(ctx *gin.Context)
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
	Getvalvestatus(ctx *gin.Context) int
	GetBaterai(ctx *gin.Context) float64
}

type deviceController struct {
	service service.DeviceService
}

func New(service service.DeviceService) DeviceController {
	return &deviceController{
		service: service,
	}
}

func (c *deviceController) FindAll() []entity.Device {
	return c.service.FindAll()
}

func (c *deviceController) Save(ctx *gin.Context) error {
	var device entity.Device
	err := ctx.ShouldBindJSON(&device)
	if err != nil {
		return err
	}
	c.service.Save(device)
	return nil
}
func (c *deviceController) Show(ctx *gin.Context) {
	device := c.service.FindAll()
	ctx.HTML(http.StatusOK, "index.html", device)
}
func (c *deviceController) Update(ctx *gin.Context) error {
	var device entity.Device
	err := ctx.ShouldBindJSON(&device)
	merge_id := string(ctx.Param("merge_id"))
	if merge_id == device.Merge_id {
		device.Merge_id = merge_id
	} else {
		return err
	}

	c.service.Update(device)
	return nil
}
func (c *deviceController) Delete(ctx *gin.Context) error {
	var device entity.Device
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	device.Id = id
	c.service.Delete(device)
	return nil
}
func (c *deviceController) Getvalvestatus(ctx *gin.Context) int {
	var device entity.Device
	device.Merge_id = string(ctx.Param("merge_id"))
	hehe := c.service.Getvalvestatus(device)
	return hehe
}
func (c *deviceController) GetBaterai(ctx *gin.Context) float64 {
	var device entity.Device
	device.Merge_id = string(ctx.Param("merge_id"))
	hehe := c.service.GetBaterai(device)
	return hehe
}
