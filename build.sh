#!/bin/bash

export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0
export GOPATH=${GOPATH}:$(pwd)

env=$2
server_name=samh-manhua
server_path=src/server/main/main.go
server_name_test=${server_name}"-test"
server_path_test=src/process/main/main.go
pack_path=bin/${server_name}.tar.gz

if [[ $1 == "b" ]]
then
    go build -o bin/${server_name} ${server_path}
    go build -o bin/${server_name_test} ${server_path_test}
    exit
elif [[ $1 == "t" ]]
then
    ./bin/$server_name_test -env ${env}
elif [[ $1 == "r" ]]
then
    ./bin/${server_name} -env ${env}
elif [[ $1 == "p" ]]
then
    mkdir -p ${server_name}/bin
    cp bin/${server_name} bin/${server_name_test} ${server_name}/bin
    cp -r config run.sh ${server_name}
    tar -zcvf ${pack_path} ${server_name}
    rm -rf ${server_name}
    exit
elif [[ $1 == "s" ]]
then
    scp ${pack_path} root@47.99.39.114:/home/lingyuanwei/xndm_work/server/pack
    scp run2.sh root@47.99.39.114:/home/lingyuanwei/xndm_work/server
    exit
elif [[ $1 == "st" ]]
then
    scp ${pack_path} root@118.31.21.181:/home/lingyuanwei/xndm_work/server/pack
    scp run2.sh root@118.31.21.181:/home/lingyuanwei/xndm_work/server
    exit
elif [[ $1 == "pt" ]]
then
    mkdir -p ${server_name}/bin
    cp bin/${server_name} bin/${server_name_test} ${server_name}/bin
    cp -r config ${server_name}
    now_time=$(date "+%Y-%m-%d_%H:%M:%S")
    tar -zcvf bin/${server_name}"."${now_time}.tar.gz ${server_name}
    rm -rf ${server_name}
    exit
else
    echo "参数1 b:构建,t:本地测试,r:本地跑服务,p:打包成部署用的压缩包到bin里"
fi

if [[ $2 == "" ]]
then
    echo "参数2缺失:local,dev,test,official"
    exit
fi

if [[ $2 != "local" ]] && [[ $2 != "dev" ]] && [[ $2 != "test" ]] && [[ $2 != "official" ]]
then
    echo "参数2错误:local,dev,test,official"
    exit
fi
