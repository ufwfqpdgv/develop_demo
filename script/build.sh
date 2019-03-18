#!/bin/bash

source ./base_setting.sh

if [[ $1 == "b" ]]
then
    go build -o ../bin/${server_name} ${server_path}
    go build -o ../bin/${server_name_test} ${server_path_test}
    exit
elif [[ $1 == "t" ]]
then
    ../bin/$server_name_test -env ${env}
elif [[ $1 == "r" ]]
then
    ../bin/${server_name} -env ${env}
elif [[ $1 == "p" ]]
then
    mkdir -p ${server_name}/bin && mkdir -p ${server_name}/script
    cp ../bin/${server_name} ../bin/${server_name_test} ${server_name}/bin
    cp -r ../config ${server_name}
    cp -r  *.sh ${server_name}/script
    tar -zcvf ${pack_path} ${server_name}
    rm -rf ${server_name}
    exit
elif [[ $1 == "s" ]]
then
    scp ${pack_path} ${user_name}@${develop_addr}:${server_dir}/pack
    scp run2.sh ${user_name}@${develop_addr}:${server_dir}
    exit
elif [[ $1 == "st" ]]
then
    scp ${pack_path} ${user_name}@${test_addr}:${server_dir}/pack
    scp run2.sh ${user_name}@${test_addr}:${server_dir}
    exit
elif [[ $1 == "st2" ]]
then
    scp ${pack_path} ${user_name2}@${test_addr2}:${server_dir}/pack
    scp run2.sh ${user_name2}@${test_addr2}:${server_dir}
    exit
elif [[ $1 == "pt" ]]
then
    mkdir -p ${server_name}/bin && mkdir -p ${server_name}/script
    cp ../bin/${server_name} ../bin/${server_name_test} ${server_name}/bin
    cp -r ../config ${server_name}
    cp -r  *.sh ${server_name}/script
    now_time=$(date "+%Y-%m-%d_%H:%M:%S")
    tar -zcvf ../bin/${server_name}"."${now_time}.tar.gz ${server_name}
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
