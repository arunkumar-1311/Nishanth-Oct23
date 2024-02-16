package service

import "github.com/google/uuid"

type UUID interface {
	UniqueID() string
}

func (Service) UniqueID() string {
	return uuid.NewString()
}
