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
