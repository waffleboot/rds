#!/bin/bash

docker run -it --network some-network --rm redis redis-cli -h replica

# docker exec -it some-redis bash
