#!/usr/bin/env bash



if [[ ! -d "artifacts" ]]; then
  mkdir artifacts
else
   rm -rf artifacts/*
fi

#if [[ ! -d "crypto-config" ]]; then
#    mkdir crypto-config
#else
#    rm -rf crypto-config/*
#fi

export FABRIC_CFG_PATH=${PWD}

#cryptogen generate --config=crypto-config.yaml
fabric-ca-client enroll -d -u https://peer1.vlink.link:peer1pw@0.0.0.0:7052 --enrollment.profile tls --csr.hosts peer1.vlink.link -M ${GOPATH}/src/vlink.link/crypto-config/peerOrganizations/vlink.link/peers/peer1.vlink.link/tls --tls.certfiles ${GOPATH}/src/vlink.link/crypto-config/peerOrganizations/vlink.link/tlsca/ca-cert.pem

if [[ $? -ne 0 ]]; then
    echo "生成证书失败"
    exit -1
fi

echo "初始化创世块"
configtxgen  --profile SoloOrdererGenesis -channelID sysdemochannel -outputBlock ./artifacts/orderer.genesis.block
echo "生成channel的配置信息"
configtxgen  --profile DemoChannel  -outputCreateChannelTx ./artifacts/demochannel.tx -channelID demochannel
echo "生成组织1的锚节点信息"
configtxgen  --profile DemoChannel -outputAnchorPeersUpdate ./artifacts/demomspanchors.tx -channelID demochannel -asOrg DemoOrgMSP