package controller

import (
	"pampam/backend/tuqa/entity"
	"pampam/backend/tuqa/service"

	"github.com/gin-gonic/gin"
)

type HardwareController interface {
	Updatebabang(ctx *gin.Context) int
}

type hardwareController struct {
	service service.HardwareService
}

func NewHardwareController(service service.HardwareService) HardwareController {
	return &hardwareController{
		service: service,
	}
}

func (c *hardwareController) Updatebabang(ctx *gin.Context) int {
	var volume entity.Volume
	var device entity.Device
	mergeid := string(ctx.Param("merge_id"))
	volumes := string(ctx.Param("volume"))
	baterai := string(ctx.Param("baterai"))
	device.Merge_id = mergeid
	volume.Volume = volumes
	volume.Device_index = mergeid
	device.Indikator_baterai = baterai
	c.service.Updatebabang(volume, device)
	return 1
}
