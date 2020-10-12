#!/usr/bin/env bash

if [[ ! -d "artifacts" ]]; then
  mkdir artifacts
else
   rm -rf artifacts/*
fi

export FABRIC_CFG_PATH=${GOPATH}/src/${DIRECTORY_NAME}



if [[ $? -ne 0 ]]; then
    echo "生成证书失败"
    exit -1
fi

#echo "solo 创世快"
#configtxgen  --profile OrdererSoloGenesis -channelID sysdemochannel -outputBlock ./artifacts/orderer.solo.genesis.block
echo "初始化创世块"
configtxgen  --profile OrdererRaftGenesis -channelID sysdemochannel -outputBlock ./artifacts/orderer.genesis.block
echo "生成channel的配置信息"
configtxgen  --profile DemoChannel  -outputCreateChannelTx ./artifacts/demochannel.tx -channelID demochannel
echo "生成组织1的锚节点信息"
configtxgen  --profile DemoChannel -outputAnchorPeersUpdate ./artifacts/demomspanchors.tx -channelID demochannel -asOrg Org0MSP