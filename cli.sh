#!/bin/bash

# docker run -it --network redis-network --rm redis redis-cli -h sentinel-1 -p 5000

# docker run -it --network some-network --rm redis redis-cli -h some-redis

# docker exec -it sentinel-1 redis-cli -p 5000

# docker run -it --net=host --rm redis redis-cli -p 16379

docker exec -it redis redis-cli -p 6379
