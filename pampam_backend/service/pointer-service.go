package service

import (
	"pampam/backend/tuqa/entity"
	"pampam/backend/tuqa/repository"
)

type PointerService interface {
	FindMerge(user entity.User) []entity.Pointer
}

type pointerService struct {
	pointerrepo repository.PointerRepo
}

func NewPointerService(repo repository.PointerRepo) PointerService {
	return &pointerService{
		pointerrepo: repo,
	}
}

func (service *pointerService) FindMerge(user entity.User) []entity.Pointer {
	return service.pointerrepo.FindMerge(user)
}
