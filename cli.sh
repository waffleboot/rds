#!/bin/bash

docker run -it --network redis-network --rm redis redis-cli -h sentinel-1 -p 5000

# docker run -it --network some-network --rm redis redis-cli -h some-redis

# docker exec -it some-redis bash
