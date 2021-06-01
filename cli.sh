#!/bin/bash

docker run -it --network some-network --rm redis redis-cli -h some-redis

