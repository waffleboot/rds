#!/bin/bash

docker run -it --network some-network --rm redis redis-cli --cluster create 10.5.0.5:6379 10.5.0.6:6379 10.5.0.7:6379
