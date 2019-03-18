#!/bin/bash

export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0
export GOPATH=`cd .. && pwd`:${GOPATH}

short_server_name=comic
server_name=samh-${short_server_name}
log_location="/data1/logs/samh/${short_server_name}/log.log"
server_dir="/home/lingyuanwei/xndm_work/server"
develop_addr="47.99.39.114"
test_addr="118.31.21.181"
test_addr2="116.62.56.235"
user_name="root"
user_name2="lingyuanwei"
env=$2
server_path=../src/server/main/main.go
server_name_test=${server_name}"-test"
server_path_test=../src/process/main/main.go
pack_path=../bin/${server_name}.tar.gz
