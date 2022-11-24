package db

import (
	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

type Engine = xorm.Engine

func New(isDevelopment bool) (*Engine, error) {
	engine, err := xorm.NewEngine("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}

	engine.SetMapper(names.GonicMapper{})
	if isDevelopment {
		engine.ShowSQL(true)
	}

	return engine, nil
}
