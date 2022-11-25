package model

type Message struct {
	ID     `xorm:"text pk"`
	UserID ID `xorm:"index"`
	Body   string
}
