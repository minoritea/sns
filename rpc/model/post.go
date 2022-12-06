package model

type PostID string

func (p PostID) String() string { return string(p) }
func NewPostID() PostID         { return NewID[PostID]() }

type Post struct {
	ID     PostID `xorm:"text pk not null"`
	UserID `xorm:"text index not null"`
	Body   string `xorm:"not null"`
}
