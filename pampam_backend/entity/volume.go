package entity

import "time"

type Volume struct {
	Id           uint64    `json:"id" gorm:"primary_key;auto_increment"`
	Volume       string    `json:"volume"`
	Device_index string    `json:"device_index"`
	Created_at   time.Time `json:"created_at"`
}
