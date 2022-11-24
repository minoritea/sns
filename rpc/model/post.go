package model

type Post struct {
	ID   `xorm:"text pk"`
	Body string
}
