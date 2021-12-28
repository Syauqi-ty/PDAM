package service

import (
	"pampam/backend/tuqa/entity"
	"pampam/backend/tuqa/repository"
)

type DeviceService interface {
	Save(entity.Device) entity.Device
	FindAll() []entity.Device
	Update(device entity.Device)
	Delete(device entity.Device)
	Getvalvestatus(device entity.Device) int
	GetBaterai(device entity.Device) float64
}

type deviceService struct {
	devicerepo repository.DeviceRepo
}

func New(repo repository.DeviceRepo) DeviceService {
	return &deviceService{
		devicerepo: repo,
	}
}

func (service *deviceService) Save(alat entity.Device) entity.Device {
	service.devicerepo.Save(alat)
	return alat
}

func (service *deviceService) FindAll() []entity.Device {
	return service.devicerepo.FindAll()
}
func (service *deviceService) Update(device entity.Device) {
	service.devicerepo.Update(device)
}
func (service *deviceService) Delete(device entity.Device) {
	service.devicerepo.Delete(device)
}
func (service *deviceService) Getvalvestatus(device entity.Device) int {
	return service.devicerepo.Getvalvestatus(device)
}
func (service *deviceService) GetBaterai(device entity.Device) float64 {
	return service.devicerepo.GetBaterai(device)
}
