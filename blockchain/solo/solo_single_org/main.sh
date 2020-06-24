#!/usr/bin/env bash

docker-compose -f docker-compose-orderer.yaml -f docker-compose-peer.yaml -f docker-compose-cli.yaml --project-name containers up -d

echo "添加到channel中" && \

peer channel create -o orderer0.demo.com:7050 -c demochannel -f ./artifacts/demo.tx --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/demo.com/orderers/orderer0.demo.com/msp/tlscacerts/tlsca.demo.com-cert.pem && \

echo "将该组织加入到该channel中" && \
peer channel join -b demochannel.block && \


echo "将其他组织也加入到peer中" && \
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.demo.com/users/Admin@org2.demo.com/msp CORE_PEER_ADDRESS=peer0.org2.demo.com:9051 CORE_PEER_LOCALMSPID=Org2MSP CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.demo.com/peers/peer0.org2.demo.com/tls/ca.crt peer channel join -b demochannel.block && \

echo "更新当前组织的anchor peer" && \
peer channel update -o orderer0.demo.com:7050 -c demochannel -f ./artifacts/org1mspanchors.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/demo.com/orderers/orderer0.demo.com/msp/tlscacerts/tlsca.demo.com-cert.pem && \
echo "更新组织2的anchor peer" && \
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.demo.com/users/Admin@org2.demo.com/msp CORE_PEER_ADDRESS=peer0.org2.demo.com:9051 CORE_PEER_LOCALMSPID=Org2MSP CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.demo.com/peers/peer0.org2.demo.com/tls/ca.crt peer channel update -o orderer0.demo.com:7050 -c demochannel -f ./artifacts/org2mspanchors.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/demo.com/orderers/orderer0.demo.com/msp/tlscacerts/tlsca.demo.com-cert.pem && \

echo "安装chaincode" && \
peer chaincode install -n democc -v 1.0 -p github.com/chaincode && \

echo "在其他节点中也安装chaincode" && \
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.demo.com/users/Admin@org2.demo.com/msp CORE_PEER_ADDRESS=peer0.org2.demo.com:9051 CORE_PEER_LOCALMSPID=Org2MSP CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.demo.com/peers/peer0.org2.demo.com/tls/ca.crt peer chaincode install -n democc -v 1.0 -p github.com/chaincode && \
echo "实例化chaincode" && \
peer chaincode instantiate -o orderer0.demo.com:7050 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/demo.com/orderers/orderer0.demo.com/msp/tlscacerts/tlsca.demo.com-cert.pem -C demochannel -n democc -v 1.0 -c '{"Args":["init","a","b","100"]}' -P "AND ('Org1MSP.peer','Org2MSP.peer')" && \
echo "查询测试" && \
peer chaincode query -C demochannel -n democc -c '{"Args":["query","a"]}'