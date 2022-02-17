package models

type User struct {
	Id       string `bson:"_id,omitempty" json:"id,omitempty"`
	Username string `json:"username"`
}
