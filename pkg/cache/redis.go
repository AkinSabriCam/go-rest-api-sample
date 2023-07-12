package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type Redis struct {
	client *redis.Client
}

// NewRedis method creates a new redis client
func NewRedis() *Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set (default)
		DB:       0,  // use default DB
	})

	if client == nil {
		fmt.Println("error could not connect to redis")
	}

	return &Redis{
		client: client,
	}

}

// Get method gets the value from redis
func (r Redis) Get() []byte {
	strCmd := r.client.Get(ctx, "products")

	cacheBytes, err := strCmd.Bytes()

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return cacheBytes
}

// Put method puts the value to redis
func (r Redis) Put(value []byte) {
	err := r.client.Set(ctx, "products", value, 0).Err()
	if err != nil {
		panic(err)
	}
}
