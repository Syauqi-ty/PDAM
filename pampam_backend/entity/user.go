package entity

type User struct {
	Id       uint64 `json:"id" gorm:"primary_key;auto_increment"`
	Nama     string `json:"nama"`
	Phone    string `json:"phone"`
	NIK      string `json:"NIK"`
	Username string `json:"username"`
	Password string `json:"password"`
}
