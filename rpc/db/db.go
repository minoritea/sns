package db

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

type Engine = xorm.Engine

func New(isDevelopment bool) (*Engine, error) {
	engine, err := xorm.NewEngine("sqlite3", "sns.db")
	if err != nil {
		return nil, err
	}

	engine.SetMapper(names.GonicMapper{})
	if isDevelopment {
		engine.ShowSQL(true)
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
