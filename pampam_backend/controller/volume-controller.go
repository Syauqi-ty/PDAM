package controller

import (
	"pampam/backend/tuqa/entity"
	"pampam/backend/tuqa/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type VolumeController interface {
	FindAll() []entity.Volume
	Save(ctx *gin.Context) error
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) int
	FindID(ctx *gin.Context) hasil
	FindMonthly(ctx *gin.Context) hasil
	DataTerakhir(ctx *gin.Context) entity.Last
	ArrayDaily(ctx *gin.Context) []entity.Array
	ArrayHourly(ctx *gin.Context) []entity.Array
}

type volumeController struct {
	service service.VolumeService
}

type hasil struct {
	Volume  float64
	Message string
}

func NewVolumeController(service service.VolumeService) VolumeController {
	return &volumeController{
		service: service,
	}
}

func (c *volumeController) FindAll() []entity.Volume {
	return c.service.FindAll()
}

func (c *volumeController) Save(ctx *gin.Context) error {
	var volume entity.Volume
	err := ctx.ShouldBindJSON(&volume)
	if err != nil {
		return err
	}
	c.service.Save(volume)
	return nil
}

func (c *volumeController) Update(ctx *gin.Context) error {
	var volume entity.Volume
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	volume.Id = id
	c.service.Delete(volume)
	return nil
}
func (c *volumeController) Delete(ctx *gin.Context) int {
	var volume entity.Volume
	id := string(ctx.Param("id"))
	volume.Device_index = id
	c.service.Delete(volume)
	return 1
}
func (c *volumeController) FindID(ctx *gin.Context) hasil {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)
	var device entity.Device
	var lempar hasil
	device.Merge_id = string(ctx.Param("merge_id"))
	_, hasil := c.service.FindID(device)
	lempar.Volume = hasil
	msghar := now.Day()
	pesan := strconv.Itoa(msghar)
	lempar.Message = "Query berhasil untuk daily tanggal " + pesan
	return lempar
}

func (c *volumeController) FindMonthly(ctx *gin.Context) hasil {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)
	var device entity.Device
	var lempar hasil
	device.Merge_id = string(ctx.Param("merge_id"))
	_, hasil := c.service.FindMonthly(device)
	lempar.Volume = hasil
	msghar := now.Month()
	lempar.Message = "Query berhasil untuk bulan " + msghar.String()
	return lempar
}

func (c *volumeController) DataTerakhir(ctx *gin.Context) entity.Last {
	var volume entity.Volume
	volume.Device_index = string(ctx.Param("device_index"))
	pantat := c.service.DataTerakhir(volume)
	return pantat
}
func (c *volumeController) ArrayDaily(ctx *gin.Context) []entity.Array {
	var device entity.Device
	device.Merge_id = string(ctx.Param("merge_id"))
	pantat := c.service.ArrayDaily(device)
	return pantat
}
func (c *volumeController) ArrayHourly(ctx *gin.Context) []entity.Array {
	var device entity.Device
	device.Merge_id = string(ctx.Param("merge_id"))
	pantat := c.service.ArrayHourly(device)
	return pantat
}
