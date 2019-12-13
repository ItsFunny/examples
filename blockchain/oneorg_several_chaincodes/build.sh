#!/usr/bin/env bash
echo "clean out-of-date datas"
rm -rf ./artifacts/*
rm -rf ./crypto-config/*
cryptogen generate --config=crypto-config.yaml
configtxgen --profile OneOrgOrdererSoloGenesis -outputBlock ./artifacts/orderer.genesis.block
configtxgen --profile OneOrgChannel -outputCreateChannelTx ./artifacts/demo.tx -channelID demo
configtxgen --profile OneOrgChannel -outputAnchorPeersUpdate ./artifacts/org1.tx -channelID demo -asOrg Org1MSP
