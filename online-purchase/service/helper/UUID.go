package helper

import "github.com/google/uuid"

func UniqueID() string {
	return uuid.NewString()
}
