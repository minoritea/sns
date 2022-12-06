package model

type UserID string

func (u UserID) String() string { return string(u) }
func NewUserID() UserID         { return NewID[UserID]() }

type User struct {
	ID       UserID `xorm:"text pk not null"`
	Name     string `xorm:"unique not null"`
	Email    string `xorm:"unique not null"`
	Password string `xorm:"not null"`
}

const PasswordCost = 10
