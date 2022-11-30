package model

type User struct {
	ID       `xorm:"text pk not null"`
	Name     string `xorm:"unique not null"`
	Email    string `xorm:"unique not null"`
	Password string `xorm:"not null"`
}
