#!/bin/bash

docker run -it --network some-network --rm redis redis-cli -c -h redis-a

# docker run -it --network some-network --rm redis redis-cli -h some-redis

# docker exec -it some-redis bash
