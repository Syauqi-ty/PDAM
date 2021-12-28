package entity

type Device struct {
	Id                uint64 `json:"id" gorm:"primary_key;auto_increment"`
	Merge_id          string `json:"merge_id"`
	Valve_status      int    `json:"valve_status"`
	Indikator_baterai string `json:"indikator_baterai"`
	Device_id         string `json:"device_id"`
	Gateway_id        string `json:"gateway_id"`
}
