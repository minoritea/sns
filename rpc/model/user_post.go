package model

type UserPost struct {
	*User `xorm:"extends"`
	*Post `xorm:"extends"`
}
