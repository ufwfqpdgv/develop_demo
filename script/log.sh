#!/bin/bash

source ./base_setting.sh

if [[ $1 == "l" ]]
then
    tail -f ${log_location} $2
    exit
elif [[ $1 == "d" ]]
then
    ssh -t ${user_name}@${develop_addr} "tail -f ${log_location}" $2
    exit
elif [[ $1 == "t" ]]
then
    ssh -t ${user_name}@${test_addr} "tail -f ${log_location}" $2
    exit
elif [[ $1 == "ol" ]]
then
    vim ${log_location}
    exit
elif [[ $1 == "od" ]]
then
    vim scp://${user_name}@${develop_addr}/${log_location}
    exit
elif [[ $1 == "ot" ]]
then
    vim scp://${user_name}@${test_addr}/${log_location}
    exit
fi
