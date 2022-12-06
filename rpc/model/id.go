package model

import "github.com/oklog/ulid/v2"

type Identifiable interface {
	~string
	String() string
}

func NewID[T Identifiable]() T { return T(ulid.Make().String()) }
