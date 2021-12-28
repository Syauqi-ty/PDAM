package service

import (
	"pampam/backend/tuqa/entity"
	"pampam/backend/tuqa/repository"
)

type AlamatService interface {
	Save(entity.Alamat) entity.Alamat
	Update(alamat entity.Alamat)
	Delete(alamat entity.Alamat)
	FindAlamat(alamat entity.Alamat) []entity.Alamat
}

type alamatService struct {
	alamatrepo repository.AlamatRepo
}

func NewAlamatService(repo repository.AlamatRepo) AlamatService {
	return &alamatService{
		alamatrepo: repo,
	}
}

func (service *alamatService) Save(alamat entity.Alamat) entity.Alamat {
	service.alamatrepo.Save(alamat)
	return alamat
}

func (service *alamatService) FindAlamat(alamat entity.Alamat) []entity.Alamat {
	return service.alamatrepo.FindAlamat(alamat)
}
func (service *alamatService) Update(alamat entity.Alamat) {
	service.alamatrepo.Update(alamat)
}
func (service *alamatService) Delete(alamat entity.Alamat) {
	service.alamatrepo.Delete(alamat)
}
