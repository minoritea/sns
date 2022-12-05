package db

import (
	"fmt"

	"github.com/minoritea/sns/rpc/model"
)

func GetUser[ID model.UserID | string](db *Engine, id ID) (*model.User, error) {
	var user model.User
	err := MustOne(db.ID(id).Get(&user))
	if err != nil {
		return nil, fmt.Errorf("GetUser failed(id: %s): %w", id, err)
	}
	return &user, nil
}

type UserParameter interface {
	GetName() string
	GetEmail() string
	GetPassword() string
}

func CreateUser(db *Engine, u UserParameter) (*model.User, error) {
	user := model.User{ID: model.NewUserID(), Name: u.GetName(), Email: u.GetEmail(), Password: u.GetPassword()}
	_, err := db.Insert(&user)
	return &user, err
}
