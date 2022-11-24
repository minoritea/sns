package model

import "github.com/oklog/ulid/v2"

type ID string

func NewID() ID              { return ID(ulid.Make().String()) }
func (id ID) String() string { return string(id) }
