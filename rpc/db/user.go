package db

import (
	"github.com/minoritea/sns/rpc/model"
	"golang.org/x/crypto/bcrypt"
)

type UserParameter interface {
	GetName() string
	GetEmail() string
	GetPassword() string
}

func CreateUser(db DB, u UserParameter) (*model.User, error) {
	pw, err := bcrypt.GenerateFromPassword([]byte(u.GetPassword()), model.PasswordCost)
	if err != nil {
		return nil, err
	}
	user := model.User{ID: model.NewUserID(), Name: u.GetName(), Email: u.GetEmail(), Password: string(pw)}
	_, err = db.Insert(&user)
	return &user, err
}

func FindUserBySessionID(db DB, sessionID model.SessionID) (*model.User, error) {
	var userSession model.UserSession
	err := MustOne(
		db.
			Table("user").
			Join("inner", "session", "user.id = session.user_id").
			Where("session.id = ?", sessionID).
			Get(&userSession),
	)
	return userSession.User, err
}

func FindUserByName(db DB, name string) (*model.User, error) {
	var user model.User
	err := MustOne(db.Where("name = ?", name).Get(&user))
	return &user, err
}
