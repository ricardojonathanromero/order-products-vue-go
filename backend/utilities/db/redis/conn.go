package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

func New(addr, pass string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
	})

	status := client.Ping(context.Background())
	if status.Err() != nil {
		return client, status.Err()
	}

	return client, nil
}
