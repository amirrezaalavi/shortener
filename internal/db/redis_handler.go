package database

import (
	"context"
	"fmt"
	"sync"
	"time"

	redis "github.com/redis/go-redis/v9"
)

type Entity struct {
	//Id uint
	ShortSuffix string
	Url         string
	TTL         uint64
}

var (
	DBClient     *redis.Client
	DBClientOnce sync.Once
	DBClientErr  error

	CTX = context.Background()
)

func create_connection_database() *redis.Client {
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
func get_url_database(rdb *redis.Client, Suffix string) (url string) {
	url, err := rdb.HGet(CTX, Suffix, "URL").Result()
	if err != nil {
		panic(err)
	}
	return url

}

func exists_Suffix_database(rdb *redis.Client, Suffix string) (exists bool) { // Check if a Suffix exists by querying the database
	output := rdb.HGet(CTX, Suffix, "Created_Time")
	if output.Err() != nil {
		fmt.Println("Entity Exists")
		return true
	}
	return false
}
func create_entity_database(rdb *redis.Client, Url string, Suffix string) (err error) {
	output := rdb.HSet(CTX, Suffix, "URL", Url, "Suffix", Suffix, "Created_Time", time.Time.Unix(time.Now())) // I'll use Suffix as a hash identifier for easier deduplication
	err = output.Err()
	return err

}

func close_database(rdb *redis.Client) {
	rdb.Close()
}
