package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/go-redis/redis/v8"

	_ "net/http/pprof"

	_ "net/http"
)

type MyLogger struct {
}

func (_ MyLogger) Printf(ctx context.Context, format string, v ...interface{}) {
}

func test(ctx context.Context, rdb *redis.Client) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			fmt.Println(rdb.Ping(ctx))
		case <-ctx.Done():
			return
		}
	}
}

func main() {

	// redis.SetLogger(MyLogger{})
	rdb := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "mymaster",
		SentinelAddrs: []string{"10.5.0.5:5000", "10.5.0.6:5000", "10.5.0.7:5000"},
	})
	defer rdb.Close()

	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	defer close(c)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		signal.Stop(c)
		cancel()
	}()
	test(ctx, rdb)
}
