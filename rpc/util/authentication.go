package util

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/minoritea/sns/rpc/model"
)

func Authenticate(user *model.User, password string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	)
}
