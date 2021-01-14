#!/usr/bin/env bash

echo "安装chaincode"
for i in {1..5}
do
    echo "组织${i}安装chaincode中"
    export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org${i}.com/users/Admin@org${i}.com/msp  CORE_PEER_ADDRESS=peer0.org${i}.com:1${i}051 CORE_PEER_LOCALMSPID=Org${i}MSP CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org${i}.com/peers/peer0.org${i}.com/tls/ca.crt
    peer chaincode install -n democc -v 1.0 -p github.com/hyperledger/fabric/chaincode/
done

./scripts/instantitatechaincode.sh