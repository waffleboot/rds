package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"

	_ "net/http/pprof"

	_ "net/http"
)

type MyLogger struct {
}

func (_ MyLogger) Printf(ctx context.Context, format string, v ...interface{}) {
}

func isBusyGroup(err error) bool {
	return strings.HasPrefix(err.Error(), "BUSYGROUP")
}

func sleep(ctx context.Context, ticker *time.Ticker) bool {
	ticker.Reset(3 * time.Second)
	select {
	case <-ticker.C:
		return false
	case <-ctx.Done():
		return true
	}
}

func read(ctx context.Context, rdb *redis.Client, consumer string, last bool) error {
	ticker := time.NewTicker(3 * time.Second)
	p := "0"
	if last {
		p = ">"
	}
	for {
		s, err := rdb.XReadGroup(ctx, &redis.XReadGroupArgs{
			Group:    "mygroup",
			Consumer: consumer,
			Streams:  []string{"mystream", p},
			Count:    1,
			Block:    0,
			NoAck:    false,
		}).Result()
		if err != nil && err != redis.Nil {
			return err
		}
		for i := range s {
			if msgs := s[i].Messages; len(msgs) > 0 {
				for j := range msgs {
					fmt.Println(msgs[j].Values["name"])
					if sleep(ctx, ticker) {
						return nil
					}
					if !last {
						rdb.XAck(ctx, "mystream", "mygroup", msgs[j].ID)
						p = msgs[j].ID
					}
				}
			} else {
				return nil
			}
		}
	}
}

func run(ctx context.Context, rdb *redis.Client, consumer string) error {
	if err := rdb.XGroupCreateMkStream(ctx, "mystream", "mygroup", "$").Err(); err != nil && !isBusyGroup(err) {
		return err
	}
	if err := read(ctx, rdb, consumer, false); err != nil {
		return err
	}
	fmt.Println(">")
	if err := read(ctx, rdb, consumer, true); err != nil {
		return err
	}
	return nil
}

func getConsumer() (string, error) {
	if len(os.Args) < 2 {
		return "", errors.New("no consumer")
	}
	return os.Args[1], nil
}

func getIntSignal() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		c := make(chan os.Signal, 1)
		defer close(c)
		signal.Notify(c, os.Interrupt)
		defer signal.Stop(c)
		<-c
		cancel()
	}()
	return ctx
}

func main() {

	consumer, err := getConsumer()
	if err != nil {
		fmt.Println(err)
		return
	}

	// redis.SetLogger(MyLogger{})

	rdb := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})
	defer rdb.Close()

	if err := run(getIntSignal(), rdb, consumer); err != nil {
		fmt.Println(err)
	}

}
