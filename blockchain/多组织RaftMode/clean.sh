#!/usr/bin/env bash

if [[ ! -d "artifacts" ]]; then
  mkdir artifacts
else
   rm -rf artifacts/*
fi

if [[ ! -d "crypto-config" ]]; then
    mkdir crypto-config
else
    rm -rf crypto-config/*
fi

export FABRIC_CFG_PATH=${PWD}

cryptogen generate --config=crypto-config.yaml

if [[ $? -ne 0 ]]; then
    echo "生成证书失败"
    exit -1
fi

echo "初始化创世块"
configtxgen  --profile DemoOrdererRaftGenesis -channelID sysdemochannel -outputBlock ./artifacts/orderer.genesis.block
echo "生成channel的配置信息"
configtxgen  --profile DemoChannel  -outputCreateChannelTx ./artifacts/demo.tx -channelID demochannel
echo "生成组织1的锚节点信息"
configtxgen  --profile DemoChannel -outputAnchorPeersUpdate ./artifacts/org1mspanchors.tx -channelID demochannel -asOrg Org1MSP
echo "生成组织2的锚节点信息"
configtxgen --profile DemoChannel -outputAnchorPeersUpdate ./artifacts/org2mspanchors.tx -channelID demochannel -asOrg Org2MSP