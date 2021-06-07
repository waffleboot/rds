#!/bin/bash

# docker run -it --network some-network --rm redis redis-cli -h replica

docker exec -it resque redis-cli -p 6380
