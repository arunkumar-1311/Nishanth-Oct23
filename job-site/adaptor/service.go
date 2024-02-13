package adaptor

import "job-post/repository"

type Database interface {
	repository.Account
	repository.Post
	repository.Country
	repository.JobType
	repository.Comments
	repository.RedisCache
}
