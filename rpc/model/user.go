package model

type User struct {
	ID       `xorm:"text pk"`
	Name     string
	Password []byte
}
