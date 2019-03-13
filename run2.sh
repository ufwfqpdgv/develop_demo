#!/bin/bash

env=$2
server_name=samh-manhua

if [[ $1 == "b" ]]
then
    rm -rf ${server_name}
    tar -zxvf pack/${server_name}.tar.gz
    exit
elif [[ $1 == "r" ]]
then
    cd ${server_name}
    ./run.sh restart ${env}
else
    echo "参数1 b:重新部署,r:重新运行"
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
