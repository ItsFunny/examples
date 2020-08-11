#!/usr/bin/env bash


# docker.for.mac.host.internal 172.224.2.2
echo "192.168.65.2 orderer0.vlink.link" >> /etc/hosts
export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/admin/msp
peer channel create -o orderer0.demo.com:7050 -c demochannel -f ../artifacts/demochannel.tx --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/tls-msp/tlscacerts/tls-0-0-0-0-7052.pem  && \

export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/admin/msp

# 加入区块
peer channel join -b ${PWD}/demochannel.block


# install chaincode
export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/admin/msp
peer chaincode install -n democc -v 1.0 -p github.com/hyperledger/fabric/chaincode/

# 实例化chaincode
peer chaincode instantiate -C demochannel -n democc -v 1.0 -c '{"Args":["init","a","100","b","200"]}' -o orderer0.demo.com:7050 --tls  --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/tls-msp/tlscacerts/tls-0-0-0-0-7052.pem

# invoke
peer chaincode invoke -C demochannel -n democc -c '{"Args":["invoke","a","b","10"]}' --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/tls-msp/tlscacerts/tls-0-0-0-0-7052.pem


# 汇总
# 很可能是这一步有问题
# 提示 Cannot run peer because error when setting up MSP of type bccsp from directory /opt/gopath/src/github.com/hyperledger/fabric/peer/admin/msp: administrators must be declared when no admin ou classification is set
export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/admin/msp && \
peer channel create -o orderer0.demo.com:7050 -c demochannel -f ../artifacts/demochannel.tx --tls true --cafile /var/hyperledger/tls/ca.crt && \
export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/admin/msp && \
peer channel join -b ${PWD}/demochannel.block && \
export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/admin/msp && \
peer chaincode install -n democc -v 1.0 -p github.com/hyperledger/fabric/chaincode/ && \
export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/admin/msp && \
peer chaincode instantiate -C demochannel -n democc -v 1.0 -c '{"Args":["init","a","100","b","200"]}' -o orderer0.demo.com:7050 --tls  --cafile /var/hyperledger/tls/ca.crt && \
export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/admin/msp && \
peer chaincode invoke -C demochannel -n democc -c '{"Args":["invoke","a","b","10"]}' --tls --cafile /var/hyperledger/tls/ca.crt
