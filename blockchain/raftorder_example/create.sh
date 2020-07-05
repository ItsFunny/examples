#!/usr/bin/env bash


cd containers && \
docker-compose -f docker-compose-orderer.yaml  -f docker-compose-cli.yaml -f docker-compose-ca.yaml up  -d

docker exec -it cli /bin/bash

peer channel create -o orderer.demo.com:7050 -c demochannel -f ./artifacts/demo.tx --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/demo.com/orderers/orderer.demo.com/msp/tlscacerts/tlsca.demo.com-cert.pem


# 将该cli 绑定的peer的对应的组织加入channel
peer channel join -b demochannel.block



# 将其他peer节点join channel
#export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.demo.com/users/Admin@org2.demo.com/msp
#export CORE_PEER_ADDRESS=peer0.org2.demo.com:9051
#export CORE_PEER_LOCALMSPID=Org2MSP
#export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.demo.com/peers/peer0.org2.demo.com/tls/ca.crt
## 也可以直接这样:
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.demo.com/users/Admin@org2.demo.com/msp CORE_PEER_ADDRESS=peer0.org2.demo.com:9051 CORE_PEER_LOCALMSPID=Org2MSP CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.demo.com/peers/peer0.org2.demo.com/tls/ca.crt peer channel join -b demochannel.block


# 更新anchor peer
peer channel update -o orderer.demo.com:7050 -c demochannel -f ./artifacts/org1mspanchors.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/demo.com/orderers/orderer.demo.com/msp/tlscacerts/tlsca.demo.com-cert.pem

# 安装chaincode
peer chaincode install -n democc -v 1.0 -p github.com/chaincode

# 因为需要被org2 组织也使用,所以在org2组织中也install
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.demo.com/users/Admin@org2.demo.com/msp CORE_PEER_ADDRESS=peer0.org2.demo.com:9051 CORE_PEER_LOCALMSPID=Org2MSP CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.demo.com/peers/peer0.org2.demo.com/tls/ca.crt peer chaincode install -n democc -v 1.0 -p github.com/chaincode



# 实例化chaincode
peer chaincode instantiate -o orderer.demo.com:7050 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/demo.com/orderers/orderer.demo.com/msp/tlscacerts/tlsca.demo.com-cert.pem -C demochannel -n democc -v 1.0 -c '{"Args":["init","a","b","100"]}' -P "AND ('Org1MSP.peer','Org2MSP.peer')"


## 单组织的chaincode 整个过程
# 查询
peer chaincode query -C demochannel -n democc -c '{"Args":["query","a"]}'
# 将chaincode 安装在第三个peer节点


# -------------

# 上面的综合命令
peer channel create -o orderer.demo.com:7050 -c demochannel -f ./artifacts/demo.tx --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/demo.com/orderers/orderer.demo.com/msp/tlscacerts/tlsca.demo.com-cert.pem && \
peer channel join -b demochannel.block && \
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.demo.com/users/Admin@org2.demo.com/msp CORE_PEER_ADDRESS=peer0.org2.demo.com:9051 CORE_PEER_LOCALMSPID=Org2MSP CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.demo.com/peers/peer0.org2.demo.com/tls/ca.crt peer channel join -b demochannel.block && \
peer channel update -o orderer.demo.com:7050 -c demochannel -f ./artifacts/org1mspanchors.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/demo.com/orderers/orderer.demo.com/msp/tlscacerts/tlsca.demo.com-cert.pem && \
peer chaincode install -n democc -v 1.0 -p github.com/chaincode && \
peer chaincode instantiate -o orderer.demo.com:7050 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/demo.com/orderers/orderer.demo.com/msp/tlscacerts/tlsca.demo.com-cert.pem -C demochannel -n democc -v 1.0 -c '{"Args":["init","a","b","100"]}' -P "AND ('Org1MSP.peer','Org2MSP.peer')" && \
peer chaincode query -C demochannel -n democc -c '{"Args":["query","a"]}'





# ------------








########## 测试

# 测试用例: 测试 endorser policy 只指定一个组织,若第二个组织未加入,则是否能够对chaincode accessable
peer channel create -o orderer.demo.com:7050 -c demochannel -f ./artifacts/demo.tx --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/demo.com/orderers/orderer.demo.com/msp/tlscacerts/tlsca.demo.com-cert.pem && peer channel join -b demochannel.block && peer chaincode install -n democc -v 1.0 -p github.com/chaincode && peer chaincode instantiate -o orderer.demo.com:7050 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/demo.com/orderers/orderer.demo.com/msp/tlscacerts/tlsca.demo.com-cert.pem -C demochannel -n democc -v 1.0 -c '{"Args":["init","a","b","100"]}' -P "AND ('Org1MSP.peer','Org2MSP.peer')"
peer chaincode query -C demochannel -n democc -c '{"Args":["query","a"]}'
# 第二个组织的peer加入channel
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.demo.com/users/Admin@org2.demo.com/msp CORE_PEER_ADDRESS=peer0.org2.demo.com:9051 CORE_PEER_LOCALMSPID=Org2MSP CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.demo.com/peers/peer0.org2.demo.com/tls/ca.crt peer channel join -b demochannel.block
# 查询
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.demo.com/users/Admin@org2.demo.com/msp CORE_PEER_ADDRESS=peer0.org2.demo.com:9051 CORE_PEER_LOCALMSPID=Org2MSP CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.demo.com/peers/peer0.org2.demo.com/tls/ca.crt peer chaincode query -C demochannel -n democc -c '{"Args":["query","a"]}'
# 结果是不可以
## 子测试用例,测试安装了,但是未实例化是否能accessable
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.demo.com/users/Admin@org2.demo.com/msp CORE_PEER_ADDRESS=peer0.org2.demo.com:9051 CORE_PEER_LOCALMSPID=Org2MSP CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.demo.com/peers/peer0.org2.demo.com/tls/ca.crt peer chaincode install -n democc -v 1.0 -p github.com/chaincode
## 查询
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.demo.com/users/Admin@org2.demo.com/msp CORE_PEER_ADDRESS=peer0.org2.demo.com:9051 CORE_PEER_LOCALMSPID=Org2MSP CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.demo.com/peers/peer0.org2.demo.com/tls/ca.crt peer chaincode query -C demochannel -n democc -c '{"Args":["query","a"]}'
## 结果是可以,并且因为未以该peer为target 实例化一个container,所以也会额外实例化一个container