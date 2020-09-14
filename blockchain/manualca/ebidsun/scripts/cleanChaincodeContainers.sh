#!/bin/bash

cleanChaincodeContainer() {
    local CHAINCODE_NAME=$1
    echo "==================removing chaincodeName: $CHAINCODE_NAME start ===================="
    local INSTANTIATED_VERSION=$(docker exec cli peer chaincode list --instantiated -C ebidsun-alpha | grep "Name: $CHAINCODE_NAME," | awk -F, '{print $2}' | cut -d: -f2)
    INSTANTIATED_VERSION=${INSTANTIATED_VERSION:1}
    echo "instantiated version:" $INSTANTIATED_VERSION

    docker ps | awk '{print $1,$2}' | \
    while read cid_name; 
    do
            if [ "$cid_name" = "CONTAINER ID" ]; then
                continue
            fi
            cid=$(echo $cid_name | awk '{print $1}')
            name=$(echo $cid_name | awk '{print $2}')
            local chaincodeName=$(echo $name | awk -F- '{print $3}')
            local chaincodeVersion=$(echo $name | awk -F- '{print $4}')
            if [[ $chaincodeName == "" || $chaincodeVersion == "" ]]; then
                continue
            fi
            #echo "chaincodeName:$chaincodeName, chaincodeVersion:$chaincodeVersion"
            if [[ "$chaincodeName" = "$CHAINCODE_NAME" ]] && [[ "$chaincodeVersion" != "$INSTANTIATED_VERSION" ]]
            then
              echo "remove chaincode container $cid, chaincodeVersion:$chaincodeVersion"
               docker rm -f $cid 
            fi
    done
    echo "==================removing chaincodeName: $CHAINCODE_NAME end====================="
}

cleanAllChaincodeContainers() {
  local chaincodeNames=$(docker ps | awk '{print $2}' | awk -F- '{print $3}' | sort | uniq)
  for chaincode in $chaincodeNames
  do 
     cleanChaincodeContainer $chaincode 
  done
}

cleanAllChaincodeContainers
