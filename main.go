package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"

	_ "net/http/pprof"

	"net/http"
	_ "net/http"
)

var ctx = context.Background()

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	rdb := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "mymaster",
		SentinelAddrs: []string{"localhost:5001", "localhost:5002", "localhost:5003"},
	})
	defer rdb.Close()
	// c, cancel := context.WithTimeout(ctx, 3*time.Second)
	// defer cancel()
	// fmt.Println(rdb.Ping(c))
	fmt.Println(rdb.Ping(ctx))
}
