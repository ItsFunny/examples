#!/usr/bin/env bash
echo "clean out-of-date datas"
rm -rf ./artifacts/*
rm -rf ./crypto-config/*
echo "=====cryptogen generate --config=crypto-config.yaml======"
cryptogen generate --config=crypto-config.yaml
echo "初始化order块"
echo "=====configtxgen --profile DemoOrdererRaftGenesis -channelID demochannel -outputBlock ./artifacts/orderer.genesis.block======"
configtxgen --profile DemoOrdererRaftGenesis -channelID demochannel -outputBlock ./artifacts/orderer.genesis.block
echo "生成channel的配置信息"
echo "=====configtxgen --profile DemoChannel  -outputCreateChannelTx ./artifacts/demo.tx -channelID demochannel====="
configtxgen --profile DemoChannel  -outputCreateChannelTx ./artifacts/demo.tx -channelID demochannel

echo "生成组织org0的锚节点更新信息"
echo "=====configtxgen --profile DemoChannel -outputAnchorPeersUpdate ./artifacts/org0mspanchors.tx -channelID demochannel -asOrg Org0MSP====="
configtxgen --profile DemoChannel -outputAnchorPeersUpdate ./artifacts/org0mspanchors.tx -channelID demochannel -asOrg Org0MSP

#echo "更新userchannel的锚节点信息"
#configtxgen --profile VlinkUserChannel -outputAnchorPeersUpdate ./artifacts/user.tx -channelID userchannel -asOrg VlinkOrgMSP
#echo "更新coinchannel的锚节点信息"
#configtxgen --profile VlinkCoinChannel -outputAnchorPeersUpdate ./artifacts/coin.tx -channelID coinchannel -asOrg VlinkOrgMSP
#echo "更新copyrightchannel的锚节点信息"
#configtxgen --profile VlinkCoinChannel -outputAnchorPeersUpdate ./artifacts/copyright.tx -channelID copyrightchannel -asOrg VlinkOrgMSP
