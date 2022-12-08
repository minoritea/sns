package db

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
	"xorm.io/xorm/names"

	"github.com/minoritea/sns/rpc/model"
)

type Engine = xorm.Engine
type Session = xorm.Session
type DB = xorm.Interface

func New(isDevelopment bool) (*Engine, error) {
	engine, err := xorm.NewEngine("sqlite3", "sns.db")
	if err != nil {
		return nil, err
	}

	engine.SetMapper(names.GonicMapper{})
	if isDevelopment {
		engine.ShowSQL(true)
	}
	err = engine.Sync2(model.Post{}, model.User{}, model.Session{}, model.Following{})
	if err != nil {
		return nil, err
	}

	return engine, nil
}

var RecordNotFound = fmt.Errorf("record not found")

func MustOne(has bool, err error) error {
	if err != nil {
		return err
	}

	if !has {
		return RecordNotFound
	}

	return nil
}

func Transaction[T any](engine *Engine, f func(db DB) (*T, error)) (*T, error) {
	result, err := engine.Transaction(func(session *Session) (any, error) {
		return f(session)
	})
	if err != nil {
		return nil, err
	}
	return result.(*T), nil
}
