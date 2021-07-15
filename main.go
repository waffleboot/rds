package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	n, _ := strconv.ParseInt(os.Args[1], 10, 32)

	data := make([]byte, n)

	err := rdb.Set(ctx, "key", data, 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	for {
		<-time.After(3 * time.Second)
		rdb.Set(ctx, "key", "test", 0)
	}
}
