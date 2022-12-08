package model

type Following struct {
	FolloweeID UserID `xorm:"text unique(idx_followee_follower) not null"`
	FollowerID UserID `xorm:"text unique(idx_followee_follower) not null"`
}
