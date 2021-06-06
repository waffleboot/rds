#define REDISMODULE_EXPERIMENTAL_API
#include <stdlib.h>
#include <string.h>
#include "redismodule.h"

static void
filter (RedisModuleCommandFilterCtx *ctx) {
    size_t len;
    const char* cmd = RedisModule_StringPtrLen(RedisModule_CommandFilterArgGet(ctx,0),&len);
    if (!strcmp(cmd,"info") && RedisModule_CommandFilterArgsCount(ctx) == 1) {
        RedisModule_CommandFilterArgInsert(ctx,1,RedisModule_CreateStringPrintf(NULL,"stats"));
    }
}
int RedisModule_OnLoad(RedisModuleCtx *ctx, RedisModuleString **argv, int argc) {
    moduleCtx = ctx;
    setModuleCtx(ctx);
    if (RedisModule_Init(ctx,"filter",1,REDISMODULE_APIVER_1) == REDISMODULE_ERR) return REDISMODULE_ERR;
    RedisModule_RegisterCommandFilter(ctx, filter, REDISMODULE_CMDFILTER_NOSELF);
    return REDISMODULE_OK;
}