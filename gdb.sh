#!/bin/bash

docker exec -it --privileged redis gdb --command=/data/gdb.txt --pid 1

