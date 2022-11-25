package model

type UserMessage struct {
	*User    `xorm:"extends"`
	*Message `xorm:"extends"`
}
