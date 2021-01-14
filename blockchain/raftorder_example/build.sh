#!/usr/bin/env bash
export FABRIC_CFG_PATH=${PWD}
echo "clean out-of-date datas"
rm -rf ./artifacts/*
rm -rf ./crypto-config/*
echo "=====cryptogen generate --config=crypto-config.yaml======"
cryptogen generate --config=crypto-config.yaml
echo "初始化order块"
echo "=====configtxgen --profile DemoOrdererRaftGenesis -channelID sysdemochannel -outputBlock ./artifacts/orderer.genesis.block======"
configtxgen  --profile DemoOrdererRaftGenesis -channelID sysdemochannel -outputBlock ./artifacts/orderer.genesis.block
echo "生成channel的配置信息"
echo "=====configtxgen --profile DemoChannel  -outputCreateChannelTx ./artifacts/demo.tx -channelID demochannel====="
configtxgen  --profile DemoChannel  -outputCreateChannelTx ./artifacts/demo.tx -channelID demochannel

echo "生成组织org1的锚节点更新信息"
echo "=====configtxgen --profile DemoChannel -outputAnchorPeersUpdate ./artifacts/org0mspanchors.tx -channelID demochannel -asOrg Org0MSP====="
configtxgen  --profile DemoChannel -outputAnchorPeersUpdate ./artifacts/org1mspanchors.tx -channelID demochannel -asOrg Org1MSP

configtxgen --profile DemoChannel -outputAnchorPeersUpdate ./artifacts/org2mspanchors.tx -channelID demochannel -asOrg Org2MSP

