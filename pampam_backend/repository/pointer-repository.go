package repository

import (
	"pampam/backend/tuqa/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PointerRepo interface {
	FindMerge(user entity.User) []entity.Pointer
}

type databaseP struct {
	connection *gorm.DB
}

func NewPointerRepo() PointerRepo {
	dsn := "host=localhost user=postgres password=250330 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to konak")
	}
	db.AutoMigrate(&entity.Pointer{})
	return &databaseP{
		connection: db,
	}
}

func (db *databaseP) FindMerge(user entity.User) []entity.Pointer {
	var pointer []entity.Pointer
	userid := user.Id
	db.connection.Where("user_id = ?", userid).Find(&pointer)
	return pointer
}
