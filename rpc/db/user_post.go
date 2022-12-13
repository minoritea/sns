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

func FindUserAndFolloweePosts(db DB, userID model.UserID, limit int) ([]model.UserPost, error) {
	var userPosts []model.UserPost
	err := db.SQL(`
		select u.*, p.*
		from (
			select * from user where user.id = ?
			union all
			select followee.* from user as followee
			inner join following on following.followee_id = followee.id
			inner join user as follower on following.follower_id = follower.id
			where follower.id = ?
		) as u
		inner join post as p on p.user_id = u.id
		order by p.id desc
		limit ?
	`, userID, userID, limit).Find(&userPosts)
	if err != nil {
		return nil, err
	}
	return userPosts, nil
}
