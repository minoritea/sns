package model

type Message struct {
	ID     `xorm:"text pk not null"`
	UserID ID     `xorm:"text index not null"`
	Body   string `xorm:"not null"`
}
