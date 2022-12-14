package db

import (
	"github.com/minoritea/sns/rpc/model"
)

func FindUserPosts(db DB, limit int) ([]model.UserPost, error) {
	var userPosts []model.UserPost
	err := db.
		Table("user").
		Join("inner", "post", "user.id = post.user_id").
		Limit(limit).
		Desc("post.id").
		Find(&userPosts)
	if err != nil {
		return nil, err
	}
	return userPosts, nil
}
