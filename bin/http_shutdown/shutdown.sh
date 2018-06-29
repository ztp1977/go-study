#!/usr/bin/env bash

pid=$(lsof -P -n | grep 9999 | awk '{print $2}')
echo $pid

kill -TERM $pid
