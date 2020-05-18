#!/usr/bin/env bash
echo "clean out-of-date datas"
rm -rf ./artifacts/*
rm -rf ./crypto-config/*
cryptogen generate --config=crypto-config.yaml
configtxgen --profile VlinkOrdererSoloGenesis -outputBlock ./artifacts/orderer.genesis.block
echo "生成userchannel的配置信息"
configtxgen --profile VlinkUserChannel -outputCreateChannelTx ./artifacts/userchannel.tx -channelID userchannel
echo "生成coinchannel的配置信息"
configtxgen --profile VlinkCoinChannel -outputCreateChannelTx ./artifacts/coinchannel.tx -channelID coinchannel
echo "生成copyrightchannel的配置信息"
configtxgen --profile VlinkCoinChannel -outputCreateChannelTx ./artifacts/copyrightchannel.tx -channelID copyrightchannel

echo "更新userchannel的锚节点信息"
configtxgen --profile VlinkUserChannel -outputAnchorPeersUpdate ./artifacts/user.tx -channelID userchannel -asOrg VlinkOrgMSP
echo "更新coinchannel的锚节点信息"
configtxgen --profile VlinkCoinChannel -outputAnchorPeersUpdate ./artifacts/coin.tx -channelID coinchannel -asOrg VlinkOrgMSP
echo "更新copyrightchannel的锚节点信息"
configtxgen --profile VlinkCoinChannel -outputAnchorPeersUpdate ./artifacts/copyright.tx -channelID copyrightchannel -asOrg VlinkOrgMSP
