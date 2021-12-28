package repository

import (
	"pampam/backend/tuqa/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserRepo interface {
	Save(user entity.User)
	Update(user entity.User)
	Delete(user entity.User)
	FindAll() []entity.User
	Login(user entity.User) entity.User
}

type databaseU struct {
	connection *gorm.DB
}

func NewUserRepo() UserRepo {
	dsn := "host=localhost user=postgres password=250330 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to konak")
	}
	db.AutoMigrate(&entity.User{})
	return &databaseU{
		connection: db,
	}
}

func (db *databaseU) Save(user entity.User) {
	db.connection.Create(&user)
}
func (db *databaseU) Update(user entity.User) {
	db.connection.Save(&user)
}
func (db *databaseU) Delete(user entity.User) {
	db.connection.Delete(&user)
}
func (db *databaseU) FindAll() []entity.User {
	var user []entity.User
	db.connection.Set("gorm:auto_preload", true).Find(&user)
	return user
}

func (db *databaseU) Login(user entity.User) entity.User {
	var detailUser entity.User
	username := user.Username
	password := user.Password
	db.connection.Where("username = ? AND password = ?", username, password).Find(&detailUser)
	return detailUser
}
