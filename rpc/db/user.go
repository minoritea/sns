package db

import (
	"fmt"

	"github.com/minoritea/sns/rpc/model"
)

func GetUser[IDType model.ID | string](db *Engine, id IDType) (*model.User, error) {
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
	user := model.User{ID: model.NewID(), Name: u.GetName(), Email: u.GetEmail(), Password: u.GetPassword()}
	_, err := db.Insert(&user)
	return &user, err
}

type AuthenticationParameter interface {
	GetName() string
	GetPassword() string
}

func FindUserByAuthentication(db *Engine, a AuthenticationParameter) (*model.User, error) {
	var user model.User
	err := MustOne(db.Where("name = ? and password = ?", a.GetName(), a.GetPassword()).Get(&user))
	return &user, err
}
