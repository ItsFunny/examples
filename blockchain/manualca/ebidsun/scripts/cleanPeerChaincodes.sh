#!/bin/bash

# 删除一个peer节点链码容器
cleanPeerChaincodes() {
  cd /var/hyperledger/production/chaincodes/
  # 获取所有的已安装链码名称
  chaincodeNames=$(ls | awk '{print $1}' | cut -d. -f 1 | sort |uniq);
  for name in $chaincodeNames
  do  
     #获取每个链码的最新版本名称
     latest_version=$(ls | awk '{print $1}' | grep $name | cut -d. -f 2- | sort | tail -n 1)
     latest_chaincode_name=$name"."$latest_version
     echo "-------------"
     for fileName in *
     do  
        if [[ "$fileName" =~ ^"$name" ]] &&  [ "$fileName" != "$latest_chaincode_name" ] 
        then
           echo "remove file" $fileName
           # 删除非最新版本的链码
           rm -f $fileName 
        fi  
     done 
     echo "-----------------"

  done 
}


cleanPeerChaincodes
