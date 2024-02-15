package repository

import (
	"context"
	"time"
)

type Redis interface {
	SetCache(uuid, userID, name, email string) error
	GetCache(uuid string) (map[string]string, error)
	DeleteCache(uuid string) error
}

// Helps to set the userID, name and email as cache
func (d *DB_Connection) SetCache(uuid, userID, name, email string) error {

	if err := d.Redis.HMSet(context.TODO(), uuid, "userid", userID, "name", name, "email", email).Err(); err != nil {
		return err
	}
	if err := d.Redis.Expire(context.TODO(), uuid, time.Hour*24).Err(); err != nil {
		return err
	}
	return nil
}

// Helps to get all the key pair values
func (d *DB_Connection) GetCache(uuid string) (map[string]string, error) {
	cache := d.Redis.HGetAll(context.TODO(), uuid)
	return cache.Val(), cache.Err()
}

// Helps to delete the key value pair
func (d *DB_Connection) DeleteCache(uuid string) error {
	return d.Redis.Del(context.TODO(), uuid).Err()
}
