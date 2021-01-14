#!/usr/bin/env bash

if [[ ! -d artifacts ]];
then mkdir artifacts
fi
rm -rf ./artifacts/* && \
echo "初始化创世块" && \
configtxgen_1_4_4  --profile OrdererRaftGenesis -channelID sysdemochannel -outputBlock ./artifacts/orderer.genesis.block && \
echo "生成channel的配置信息" && \

if [[ ! -e ./artifacts/demochannel.tx ]];
then rm -f ./artifacts/demochannel.tx
fi

configtxgen_1_4_4  --profile DemoChannel  -outputCreateChannelTx ./artifacts/demochannel.tx -channelID demochannel


for v in {1..5}
do
configtxgen_1_4_4  --profile DemoChannel -outputAnchorPeersUpdate ./artifacts/org${v}mspanchors.tx -channelID demochannel -asOrg Org${v}MSP
done
