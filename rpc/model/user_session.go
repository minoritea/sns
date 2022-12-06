package model

type UserSession struct {
	*User    `xorm:"extends"`
	*Session `xorm:"extends"`
}
