#!/usr/bin/env bash
echo "clean out-of-date datas"
rm -rf ./artifacts/*
rm -rf ./crypto-config/*
cryptogen generate --config=crypto-config.yaml
configtxgen --profile DemoOrdererRaftGenesis -outputBlock ./artifacts/orderer.genesis.block
echo "生成channel的配置信息"
configtxgen --profile DemoChannel -outputCreateChannelTx ./artifacts/demo.tx -channelID demochannel

#echo "更新userchannel的锚节点信息"
#configtxgen --profile VlinkUserChannel -outputAnchorPeersUpdate ./artifacts/user.tx -channelID userchannel -asOrg VlinkOrgMSP
#echo "更新coinchannel的锚节点信息"
#configtxgen --profile VlinkCoinChannel -outputAnchorPeersUpdate ./artifacts/coin.tx -channelID coinchannel -asOrg VlinkOrgMSP
#echo "更新copyrightchannel的锚节点信息"
#configtxgen --profile VlinkCoinChannel -outputAnchorPeersUpdate ./artifacts/copyright.tx -channelID copyrightchannel -asOrg VlinkOrgMSP
