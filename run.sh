#!/bin/bash

export GOPATH=${GOPATH}:$(pwd)

env=$2
server_name=samh-manhua

if [[ $1 == "start" ]]
then
    nohup ./bin/${server_name} -env ${env} &
elif [[ $1 == "stop" ]]
then
    pkill ${server_name}
elif [[ $1 == "restart" ]]
then
    pkill ${server_name}
    nohup ./bin/${server_name} -env ${env} &
elif [[ $1 == "r" ]]
then
    pkill ${server_name}
    ./bin/${server_name} -env ${env}
else
    echo "参数1 start:开始,stop:停止,restart:重新运行"
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
