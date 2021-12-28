package repository

import (
	"pampam/backend/tuqa/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type AlamatRepo interface {
	Save(alamat entity.Alamat)
	Update(alamat entity.Alamat)
	Delete(alamat entity.Alamat)
	FindAlamat(alamat entity.Alamat) []entity.Alamat
}

type databaseA struct {
	connection *gorm.DB
}

func NewAlamatRepo() AlamatRepo {
	dsn := "host=localhost user=postgres password=250330 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to konak")
	}
	db.AutoMigrate(&entity.Alamat{})
	return &databaseA{
		connection: db,
	}
}

func (db *databaseA) Save(alamat entity.Alamat) {
	db.connection.Create(&alamat)
}
func (db *databaseA) Update(alamat entity.Alamat) {
	hehe := alamat.Jalan
	hihi := alamat.Kabupaten
	huhu := alamat.Kecamatan
	hoho := alamat.Kelurahan
	haha := alamat.Latitude
	hahaha := alamat.Longtitude
	db.connection.Model(&alamat).Where("merge_id = ?", alamat.Merge_id).Updates(entity.Alamat{Jalan: hehe, Kelurahan: hoho, Kecamatan: huhu, Kabupaten: hihi, Latitude: haha, Longtitude: hahaha})
}
func (db *databaseA) Delete(alamat entity.Alamat) {
	db.connection.Model(&alamat).Where("merge_id = ?", alamat.Merge_id).Delete(&alamat)
}
func (db *databaseA) FindAlamat(alamat entity.Alamat) []entity.Alamat {
	var alamatarray []entity.Alamat
	db.connection.Model(&alamat).Where("merge_id = ?", alamat.Merge_id).Find(&alamatarray)
	return alamatarray
}
