#!/usr/bin/env bash
echo "clean out-of-date datas"
rm -rf ./artifacts/*
rm -rf ./crypto-config/*
cryptogen generate --config=crypto-config.yaml
configtxgen --profile TwoOrgOrdererSoloGenesis -outputBlock ./artifacts/orderer.genesis.block
configtxgen --profile UserChannel -outputCreateChannelTx ./artifacts/userchannel.tx -channelID userchannel
configtxgen --profile CoinChannel -outputCreateChannelTx ./artifacts/coinchannel.tx -channelID coinchannel
configtxgen --profile UserChannel -outputAnchorPeersUpdate ./artifacts/user.tx -channelID userchannel -asOrg UserOrgMSP
configtxgen --profile CoinChannel -outputAnchorPeersUpdate ./artifacts/coin.tx -channelID coinchannel -asOrg CoinOrgMSP

