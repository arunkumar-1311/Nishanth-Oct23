package adaptor

import "to-do/repository"

type Database interface {
	repository.Account
	repository.Redis
	repository.Task
}
