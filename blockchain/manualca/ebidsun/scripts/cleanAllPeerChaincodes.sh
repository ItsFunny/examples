#!/bin/bash
set -eu


# 删除低版本链码容器
cleanAllPeerChaincodes() {
    # 获取所有的peer节点
    peerIds=$(docker ps | grep "hyperledger/fabric-peer" | awk '{print $1}')
    for peerId in $peerIds
    do  
       echo "peerId" $peerId
       docker exec $peerId /opt/workspace/scripts/cleanPeerChaincodes.sh  
    done 
}

cleanAllPeerChaincodes
