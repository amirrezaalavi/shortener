package database

import (
	"context"
	"fmt"
	//	"fmt"
	"sync"

	redis "github.com/redis/go-redis/v9"
)

var (
	DBClient     *redis.Client
	DBClientOnce sync.Once
	DBClientErr  error

	CTX = context.Background()
)

func create_database() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:8002",
		Password: "",
		DB:       0,
	})
	pingDB, err := rdb.Ping(CTX).Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to server:", pingDB)

	return rdb
}
func close_database(rdb *redis.Client) {
	rdb.Close()
}
