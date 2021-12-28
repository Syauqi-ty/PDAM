package repository

import (
	"pampam/backend/tuqa/entity"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DeviceRepo interface {
	Save(device entity.Device)
	Update(device entity.Device)
	Delete(device entity.Device)
	FindAll() []entity.Device
	Getvalvestatus(device entity.Device) int
	UpdateB(device entity.Device)
	GetBaterai(device entity.Device) float64
}

type database struct {
	connection *gorm.DB
}

func NewDeviceRepo() DeviceRepo {
	dsn := "host=localhost user=postgres password=250330 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to konak")
	}
	db.AutoMigrate(&entity.Device{})
	return &database{
		connection: db,
	}
}

func (db *database) Save(device entity.Device) {
	db.connection.Create(&device)
}
func (db *database) Update(device entity.Device) {
	pantat := device.Valve_status
	if pantat == 0 || pantat == 1 {
		db.connection.Model(&device).Where("merge_id = ?", device.Merge_id).Update("valve_status", pantat)
	}
}
func (db *database) Delete(device entity.Device) {
	db.connection.Delete(&device)
}
func (db *database) FindAll() []entity.Device {
	var device []entity.Device
	db.connection.Find(&device)
	return device
}
func (db *database) Getvalvestatus(device entity.Device) int {
	db.connection.Find(&device)
	hehe := device.Valve_status
	return hehe
}
func (db *database) UpdateB(device entity.Device) {
	hehe := device.Indikator_baterai
	db.connection.Model(&device).Where("merge_id = ?", device.Merge_id).Update("indikator_baterai", hehe)
}
func (db *database) GetBaterai(device entity.Device) float64 {
	db.connection.Model(&device).Where("merge_id = ?", device.Merge_id).Find(&device)
	hehe, _ := strconv.ParseFloat(device.Indikator_baterai, 64)
	return hehe
}
