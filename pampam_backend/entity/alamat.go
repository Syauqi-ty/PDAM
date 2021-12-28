package entity

type Alamat struct {
	Id         uint64 `json:"id" gorm:"primary_key;auto_increment"`
	Merge_id   string `json:"merge_id"`
	Jalan      string `json:"jalan"`
	Kelurahan  string `json:"kelurahan"`
	Kecamatan  string `json:"kecamatan"`
	Kabupaten  string `json:"kabupaten"`
	Latitude   string `json:"koordinat"`
	Longtitude string `json:"koordinat"`
}
