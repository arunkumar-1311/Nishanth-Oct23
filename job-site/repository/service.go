package repository

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type DB_Connection struct {
	DB    *gorm.DB
	Redis *redis.Client
}
