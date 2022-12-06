package model

type PostID string

func (u PostID) String() string { return string(u) }
func NewPostID() PostID         { return NewID[PostID]() }

type Post struct {
	ID     PostID `xorm:"text pk not null"`
	UserID `xorm:"text index not null"`
	Body   string `xorm:"not null"`
}
