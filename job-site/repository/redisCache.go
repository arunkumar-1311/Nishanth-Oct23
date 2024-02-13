package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisCache interface {
	SetRedisCache(UUID, token, name string) error
	GetRedisCache(UUID string) (token string, err error)
	DeleteCache(UUID ...string) error
}

// Helps to set the values as cache
func (d *DB_Connection) SetRedisCache(UUID, token, name string) error {

	if err := d.Redis.Set(context.TODO(), UUID, token, time.Hour*24); err.Err() != nil {
		return err.Err()
	}

	if err := d.Redis.Set(context.TODO(), name, UUID, time.Hour*24); err.Err() != nil {
		return err.Err()
	}
	return nil
}

// Helps to get the get the username and token from cache
func (d *DB_Connection) GetRedisCache(UUID string) (token string, err error) {

	if token, err = d.Redis.Get(context.TODO(), UUID).Result(); err != nil {
		if err == redis.Nil {
			return "", fmt.Errorf("please do login/register")
		}
		return
	}

	return
}

// Helps to delete the cache
func (d *DB_Connection) DeleteCache(UUID ...string) error {

	if err := d.Redis.Del(context.TODO(), UUID...).Err(); err != nil {
		if err == redis.Nil {
			return fmt.Errorf("please do login/register")
		}
		return err
	}

	return nil
}
