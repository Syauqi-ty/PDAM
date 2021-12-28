package controller

import (
	"pampam/backend/tuqa/entity"
	"pampam/backend/tuqa/service"

	"github.com/gin-gonic/gin"
)

type AlamatController interface {
	FindAlamat(ctx *gin.Context) []entity.Alamat
	Save(ctx *gin.Context) error
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) int
}

type alamatController struct {
	service service.AlamatService
}

func NewAlamatController(service service.AlamatService) AlamatController {
	return &alamatController{
		service: service,
	}
}

func (c *alamatController) FindAlamat(ctx *gin.Context) []entity.Alamat {
	var alamat entity.Alamat
	alamat.Merge_id = string(ctx.Param("merge_id"))
	return c.service.FindAlamat(alamat)
}

func (c *alamatController) Save(ctx *gin.Context) error {
	var alamat entity.Alamat
	err := ctx.ShouldBindJSON(&alamat)
	if err != nil {
		return err
	}
	c.service.Save(alamat)
	return nil
}

func (c *alamatController) Update(ctx *gin.Context) error {
	var alamat entity.Alamat
	id := string(ctx.Param("merge_id"))
	alamat.Merge_id = id
	err := ctx.ShouldBindJSON(&alamat)
	if err != nil {
		return err
	}

	c.service.Update(alamat)
	return nil
}
func (c *alamatController) Delete(ctx *gin.Context) int {
	var alamat entity.Alamat
	alamat.Merge_id = string(ctx.Param("merge_id"))
	c.service.Delete(alamat)
	return 1
}
