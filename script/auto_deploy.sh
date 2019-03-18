#!/bin/bash

source ./base_setting.sh

if [[ $1 == "l" ]]
then
    ./build.sh b && ./run.sh restart local && ./log.sh l $2
    exit
fi

if [[ $1 == "r" ]]
then
    ssh -t ${user_name}@${develop_addr} "cd ${server_dir} && ./run2.sh r dev&&ps -ef|grep ${server_name}" &&
    ssh -t ${user_name}@${test_addr} "cd ${server_dir} && ./run2.sh r test&&ps -ef|grep ${server_name}"
    ssh -t ${user_name2}@${test_addr2} "cd ${server_dir} && ./run2.sh r test&&ps -ef|grep ${server_name}"
    exit
fi

if [[ $1 == "d" ]]
then
    ./build.sh b&&
    ./build.sh p &&
    ./build.sh s &&
    ssh -t ${user_name}@${develop_addr} "cd ${server_dir} && ./run2.sh b&&./run2.sh r dev&&ps -ef|grep ${server_name}" &&
    ./log.sh d $2
    exit
elif [[ $1 == "t" ]]
then
    ./build.sh b&&
    ./build.sh p &&
    ./build.sh st &&
    ssh -t ${user_name}@${test_addr} "cd ${server_dir} && ./run2.sh b&&./run2.sh r test&&ps -ef|grep ${server_name}"
    ./log.sh t $2
    exit
elif [[ $1 == "t2" ]]
then
    ./build.sh b&&
    ./build.sh p &&
    ./build.sh st2 &&
    ssh -t ${user_name2}@${test_addr2} "cd ${server_dir} &&sudo ./run2.sh b&&sudo ./run2.sh r test&&ps -ef|grep ${server_name}"
    ./log.sh t2 $2
    exit
fi

./build.sh b&&
./build.sh p &&
./build.sh s &&
./build.sh st &&
ssh -t ${user_name}@${develop_addr} "cd ${server_dir} && ./run2.sh b&&./run2.sh r dev&&ps -ef|grep ${server_name}" &&
ssh -t ${user_name}@${test_addr} "cd ${server_dir} && ./run2.sh b&&./run2.sh r test&&ps -ef|grep ${server_name}" &&
ssh -t ${user_name2}@${test_addr2} "cd ${server_dir} && ./run2.sh b&&./run2.sh r test&&ps -ef|grep ${server_name}"
