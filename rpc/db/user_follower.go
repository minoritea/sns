package db

import (
	"fmt"

	"github.com/minoritea/sns/rpc/model"
)

func FindUserFollowers(db DB, userID model.UserID) ([]model.User, error) {
	var followers []model.User
	err := db.
		Table([]string{"user", "followee"}).
		Join("inner", "following", "on followee.id = following.followee_id").
		Join("inner", []string{"user", "follower"}, "on follower.id = following.follower_id").
		Where("followee.id = ?", userID).
		Select("follower.*").
		Find(&followers)
	return followers, err
}

func FollowUser(db DB, followerID, followeeID model.UserID) error {
	if followerID == followeeID {
		return fmt.Errorf("cannot follow themself")
	}
	_, err := db.Insert(model.Following{
		FolloweeID: followeeID,
		FollowerID: followerID,
	})
	return err
}
