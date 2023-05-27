#!/usr/bin/env bash
source ./path_info.cfg

begin_path=$PWD
bin_path=${begin_path}/../bin/
# for i in ${service_names[*]}; do  
for ((i = 0; i < ${#service_names[*]}; i++)); do
    echo " starting ${service_names[$i]}---------------" 
    ${bin_path}${service_names[$i]} -f ${cfg_names[$i]} &
done