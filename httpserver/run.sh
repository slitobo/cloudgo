#!/bin/bash

# 获取go version
export VERSION=$(go version 2> /dev/null | awk '{print $3}')

uname | grep Linux && make build
uname | grep Darwin && make local

./bin/httpserver
