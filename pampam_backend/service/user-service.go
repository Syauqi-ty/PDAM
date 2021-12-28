package service

import (
	"pampam/backend/tuqa/entity"
	"pampam/backend/tuqa/repository"
)

type UserService interface {
	Save(entity.User) entity.User
	FindAll() []entity.User
	Update(user entity.User)
	Delete(user entity.User)
	Login(user entity.User) entity.User
}

type userService struct {
	userrepo repository.UserRepo
}

func NewUserService(repo repository.UserRepo) UserService {
	return &userService{
		userrepo: repo,
	}
}

func (service *userService) Save(user entity.User) entity.User {
	service.userrepo.Save(user)
	return user
}

func (service *userService) FindAll() []entity.User {
	return service.userrepo.FindAll()
}
func (service *userService) Update(user entity.User) {
	service.userrepo.Update(user)
}
func (service *userService) Delete(user entity.User) {
	service.userrepo.Delete(user)
}
func (service *userService) Login(user entity.User) entity.User {
	return service.userrepo.Login(user)
}
