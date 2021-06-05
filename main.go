package main

// #define REDISMODULE_EXPERIMENTAL_API
// #include "redismodule.h"
// #include <stdlib.h>
// static void
// filter (RedisModuleCommandFilterCtx *filter) {
// }
// static int
// HelloworldRand_RedisCommand(RedisModuleCtx *ctx, RedisModuleString **argv, int argc) {
//     RedisModule_ReplyWithLongLong(ctx,rand());
//     return REDISMODULE_OK;
// }
// int RedisModule_OnLoad(RedisModuleCtx *ctx, RedisModuleString **argv, int argc) {
//     if (RedisModule_Init(ctx,"filter",1,REDISMODULE_APIVER_1) == REDISMODULE_ERR) return REDISMODULE_ERR;
//     RedisModule_RegisterCommandFilter(ctx, filter, 0);
//     if (RedisModule_CreateCommand(ctx,"helloworld.rand",
//         HelloworldRand_RedisCommand, "fast random",
//         0, 0, 0) == REDISMODULE_ERR)
//         return REDISMODULE_ERR;
//     return REDISMODULE_OK;
// }
import "C"

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	defer rdb.Close()
	for i := 1; i <= 1000; i++ {
		err := rdb.Set(ctx, fmt.Sprintf("user:%d", i), true, 1*time.Minute).Err()
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(1 * time.Second)
	}
}
