package model

import "time"

type SessionID string

func (s SessionID) String() string { return string(s) }

type Session struct {
	ID        SessionID `xorm:"text pk not null"`
	UserID    `xorm:"text index not null"`
	CreatedAt time.Time `xorm:"not null"`
}
