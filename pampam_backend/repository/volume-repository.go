package repository

import (
	"fmt"
	"pampam/backend/tuqa/entity"
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type VolumeRepo interface {
	Save(volume entity.Volume)
	Update(volume entity.Volume)
	Delete(volume entity.Volume)
	FindAll() []entity.Volume
	FindID(device entity.Device) ([]entity.Volume, float64)
	FindMonthly(device entity.Device) ([]entity.Volume, float64)
	DataTerakhir(volume entity.Volume) entity.Last
	ArrayDaily(device entity.Device) []entity.Array
	ArrayHourly(device entity.Device) []entity.Array
}

type databaseV struct {
	connection *gorm.DB
}
type Volume struct {
	Volume float64
}

func NewVolumeRepo() VolumeRepo {
	dsn := "host=localhost user=postgres password=250330 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to konak")
	}
	db.AutoMigrate(&entity.Volume{})
	return &databaseV{
		connection: db,
	}
}

func (db *databaseV) Save(volume entity.Volume) {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)
	db.connection.Create(&volume)
	db.connection.Model(&volume).Update("created_at", now)
}
func (db *databaseV) Update(volume entity.Volume) {
	db.connection.Save(&volume)
}
func (db *databaseV) Delete(volume entity.Volume) {
	db.connection.Model(&volume).Where("device_index = ?", volume.Device_index).Delete(&volume)
}
func (db *databaseV) FindAll() []entity.Volume {
	var volume []entity.Volume
	db.connection.Set("gorm:auto_preload", true).Find(&volume)
	return volume
}

func (db *databaseV) FindID(device entity.Device) ([]entity.Volume, float64) {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)
	var volume []entity.Volume
	var Count string
	var sum float64
	var arrayV []Volume
	deviceid := device.Merge_id
	y, m, d := now.Date()
	var combineshit, combineshit2 string
	combineshit = strconv.Itoa(y) + "-" + strconv.Itoa(int(m)) + "-" + strconv.Itoa(d)
	combineshit2 = strconv.Itoa(y) + "-" + strconv.Itoa(int(m)) + "-" + strconv.Itoa(d+1)
	db.connection.Table("volumes").Select("sum(CAST(volume AS DECIMAL(10,2))) as Count").Where("(device_index = ? AND created_at BETWEEN ? AND ? )", deviceid, combineshit, combineshit2).Scan(&Count)
	// x, err := strconv.ParseFloat(Count, 64)
	// if err != nil {

	// }
	db.connection.Table("volumes").Select("volume").Where("device_index = ? AND created_at BETWEEN ? AND ?", deviceid, combineshit, combineshit2).Find(&arrayV)
	for i := 0; i < len(arrayV)-1; i++ {
		a := arrayV[i].Volume
		b := arrayV[i+1].Volume
		sum = (a+b)/2 + sum
	}
	return volume, sum

}

func (db *databaseV) FindMonthly(device entity.Device) ([]entity.Volume, float64) {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)
	var volume []entity.Volume
	var Count string
	var sum float64
	var arrayV []Volume
	deviceid := device.Merge_id
	y, m, _ := now.Date()
	var combineshit, combineshit2 string
	combineshit = strconv.Itoa(y) + "-" + strconv.Itoa(int(m)) + "-1"
	combineshit2 = strconv.Itoa(y) + "-" + strconv.Itoa(int(m+1)) + "-1"
	db.connection.Table("volumes").Select("sum(CAST(volume AS DECIMAL(10,2))) as Count").Where("(device_index = ? AND created_at BETWEEN ? AND ? )", deviceid, combineshit, combineshit2).Scan(&Count)
	// x, err := strconv.ParseFloat(Count, 64)
	// if err != nil {

	// }
	db.connection.Table("volumes").Select("volume").Where("device_index = ? AND created_at BETWEEN ? AND ?", deviceid, combineshit, combineshit2).Find(&arrayV)
	for i := 0; i < len(arrayV)-1; i++ {
		a := arrayV[i].Volume
		b := arrayV[i+1].Volume
		sum = (a+b)/2 + sum
	}
	return volume, sum
}

func (db *databaseV) DataTerakhir(volume entity.Volume) entity.Last {
	db.connection.Last(&volume)
	var last entity.Last
	t := volume.Created_at
	h := t.Hour()
	m := t.Minute()
	s := t.Second()
	volumes := volume.Volume
	last.Debit, _ = strconv.ParseFloat(volumes, 64)
	last.Waktu = strconv.Itoa(h) + ":" + strconv.Itoa(m) + ":" + strconv.Itoa(s)
	return last
}

func (db *databaseV) ArrayDaily(device entity.Device) []entity.Array {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)
	var array []entity.Array
	deviceid := device.Merge_id
	d, m, y := now.Date()
	var combineshit, combineshit2 string
	combineshit = strconv.Itoa(d) + "-" + strconv.Itoa(int(m)) + "-" + strconv.Itoa(y)
	combineshit2 = strconv.Itoa(d) + "-" + strconv.Itoa(int(m)) + "-" + strconv.Itoa(y+1)
	db.connection.Table("volumes").Where("device_index = ? AND created_at BETWEEN ? AND ?", deviceid, combineshit, combineshit2).Select("volume as debit,  to_char(created_at,'HH24:MM:SS') as waktu").Find(&array)
	return array
}
func (db *databaseV) ArrayHourly(device entity.Device) []entity.Array {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)
	var array []entity.Array
	deviceid := device.Merge_id
	h, _, _ := now.Clock()
	var combineshit, combineshit2 string
	db.connection.Table("volumes").Where("device_index = ? AND EXTRACT(HOUR FROM created_at) BETWEEN ? AND ?", deviceid, h, h+1).Select("volume as debit, created_at as waktu").Find(&array)
	fmt.Print(combineshit + "-" + combineshit2)
	fmt.Print(time.Now().Date())
	return array
}
