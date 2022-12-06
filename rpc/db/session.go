package db

import (
	"time"

	"github.com/jmcvetta/randutil"

	"github.com/minoritea/sns/rpc/model"
)

func CreateSession(db DB, userID model.UserID) (*model.Session, error) {
	sessionID, err := randutil.AlphaStringRange(32, 64)
	if err != nil {
		return nil, err
	}
	session := model.Session{
		ID:        model.SessionID(sessionID),
		UserID:    userID,
		CreatedAt: time.Now(),
	}
	_, err = db.Insert(&session)
	return &session, err
}
