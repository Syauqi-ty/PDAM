package service

import (
	"pampam/backend/tuqa/entity"
	"pampam/backend/tuqa/repository"
)

type VolumeService interface {
	Save(entity.Volume) entity.Volume
	FindAll() []entity.Volume
	Update(volume entity.Volume)
	Delete(volume entity.Volume)
	FindID(device entity.Device) ([]entity.Volume, float64)
	FindMonthly(device entity.Device) ([]entity.Volume, float64)
	DataTerakhir(volume entity.Volume) entity.Last
	ArrayDaily(device entity.Device) []entity.Array
	ArrayHourly(device entity.Device) []entity.Array
}

type volumeService struct {
	volumerepo repository.VolumeRepo
}

func NewVolumeService(repo repository.VolumeRepo) VolumeService {
	return &volumeService{
		volumerepo: repo,
	}
}

func (service *volumeService) Save(volume entity.Volume) entity.Volume {
	service.volumerepo.Save(volume)
	return volume
}

func (service *volumeService) FindAll() []entity.Volume {
	return service.volumerepo.FindAll()
}
func (service *volumeService) Update(volume entity.Volume) {
	service.volumerepo.Update(volume)
}
func (service *volumeService) Delete(volume entity.Volume) {
	service.volumerepo.Delete(volume)
}
func (service *volumeService) FindID(device entity.Device) ([]entity.Volume, float64) {
	return service.volumerepo.FindID(device)
}

func (service *volumeService) FindMonthly(device entity.Device) ([]entity.Volume, float64) {
	return service.volumerepo.FindMonthly(device)
}
func (service *volumeService) DataTerakhir(volume entity.Volume) entity.Last {
	return service.volumerepo.DataTerakhir(volume)
}
func (service *volumeService) ArrayDaily(device entity.Device) []entity.Array {
	return service.volumerepo.ArrayDaily(device)
}
func (service *volumeService) ArrayHourly(device entity.Device) []entity.Array {
	return service.volumerepo.ArrayHourly(device)
}
