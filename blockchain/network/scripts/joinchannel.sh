#!/usr/bin/env bash


cd ${PWD}

sleep 5
# 组织1 加入到channel
export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.com/users/Admin@org1.com/msp && \
export ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/demo.com/orderers/orderer.demo.com/tls/ca.crt && \
peer channel create -o orderer.demo.com:5050 -c demochannel -f ../peer/artifacts/demochannel.tx --tls true --cafile ${ORDERER_CA}
sleep 5

for i in {1..5}
do
    echo "组织${i}加入到channel中"
    export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org${i}.com/users/Admin@org${i}.com/msp  CORE_PEER_ADDRESS=peer0.org${i}.com:1${i}051 CORE_PEER_LOCALMSPID=Org${i}MSP CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org${i}.com/peers/peer0.org${i}.com/tls/ca.crt
    peer channel join -b ${PWD}/demochannel.block
done

./scripts/installchaincode.sh
