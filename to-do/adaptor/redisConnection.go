package adaptor

import (
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

func RedisConnection() (*redis.Client, error) {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("host"), os.Getenv("redisport")),
		Password: "",
		DB:       1,
	})
	ctx := context.Background()
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	return client, nil
}
