#!/usr/bin/env bash
echo 'begin to build all binary files:'
source ./path_info.cfg
echo ${api_user_name}

bin_dir="../bin"
logs_dir="../logs"

#Automatically created when there is no bin, logs folder
if [ ! -d $bin_dir ]; then
  mkdir -p $bin_dir
fi
if [ ! -d $logs_dir ]; then
  mkdir -p $logs_dir
fi
#begin path
begin_path=$PWD

echo 'beginpath is: "' $begin_path '"'

for ((i = 0; i < ${#service_source_root[*]}; i++)); do
  echo $begin_path
  cd $begin_path
  service_path=${service_source_root[$i]}
  echo ${begin_path}/${service_path}
  cd $service_path
  go build -o ${begin_path}/${bin_dir}/${service_names[$i]} 
  if [ $? -ne 0 ]; then
        echo -e "${RED_PREFIX}${service_names[$i]} build failed ${COLOR_SUFFIX}\n"
        exit -1
        else
         echo -e "${GREEN_PREFIX}${service_names[$i]} successfully be built ${COLOR_SUFFIX}\n"
  fi
done
echo -e ${YELLOW_PREFIX}"all services build success"${COLOR_SUFFIX}
