package entity

type Pointer struct {
	Id       uint64 `json:id`
	Merge_id string `json:merge_id`
	User_id  uint64 `json:user_id`
}
