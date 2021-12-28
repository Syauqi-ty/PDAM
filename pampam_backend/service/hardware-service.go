package service

import (
	"pampam/backend/tuqa/entity"
	"pampam/backend/tuqa/repository"
)

type HardwareService interface {
	Updatebabang(volume entity.Volume, device entity.Device)
}

type hardwareService struct {
	devicerepo repository.DeviceRepo
	volumerepo repository.VolumeRepo
}

func NewHardwareService(repo repository.DeviceRepo, titit repository.VolumeRepo) HardwareService {
	return &hardwareService{
		devicerepo: repo,
		volumerepo: titit,
	}
}

func (service *hardwareService) Updatebabang(volume entity.Volume, device entity.Device) {
	service.devicerepo.UpdateB(device)
	service.volumerepo.Save(volume)
}
