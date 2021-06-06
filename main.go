package main

// #define REDISMODULE_EXPERIMENTAL_API
// #include "redismodule.h"
import "C"

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var moduleCtx *C.RedisModuleCtx

//export setModuleCtx
func setModuleCtx(ctx *C.RedisModuleCtx) {
	moduleCtx = ctx
}

//export getModuleCtx
func getModuleCtx() *C.RedisModuleCtx {
	return moduleCtx
}

var g func(*C.RedisModuleString, *C.size_t) *C.char = *C.RedisModule_StringPtrLen

func filter(ctx *C.struct_RedisModuleCommandFilterCtx) {
	var len C.size_t

	// REDISMODULE_API const char * (*RedisModule_StringPtrLen)(const RedisModuleString *str, size_t *len) REDISMODULE_ATTR;

	// g((*RedisModule_CommandFilterArgGet)(ctx, 0), &len)

	// const char* cmd = RedisModule_StringPtrLen(RedisModule_CommandFilterArgGet(ctx,0),&len);
	// if (!strcmp(cmd,"info") && RedisModule_CommandFilterArgsCount(ctx) == 1) {
	//     RedisModule_CommandFilterArgInsert(ctx,1,RedisModule_CreateStringPrintf(getModuleCtx(),"stats"));
	// }
}

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
