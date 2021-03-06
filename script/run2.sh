#!/bin/bash

env=$2
server_name=samh-comic

if [[ $1 == "b" ]]
then
    rm -rf ${server_name}
    tar -zxvf pack/${server_name}.tar.gz
    exit
elif [[ $1 == "r" ]]
then
    cd ${server_name}/script
    ./run.sh restart ${env}
else
    echo "参数1 b:重新部署,r:重新运行"
fi

if [[ $1 == "l" ]]
then
    cd ${server_name}/script
    ./log.sh l $2
    exit
elif [[ $1 == "ol" ]]
then
    cd ${server_name}/script
    ./log.sh ol
    exit
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
