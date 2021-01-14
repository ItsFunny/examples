#!/usr/bin/env bash


export DOMAIN=demo.com && \
export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/demo.com/users/Admin@demo.com/msp && \
peer channel create -o orderer0.${DOMAIN}:5050 -c demochannel -f ../peer/artifacts/demochannel.tx --tls true --cafile /var/hyperledger/fabric/tls/ca.crt && \
peer channel join -b ${PWD}/demochannel.block && \
peer chaincode install -n democc -v 1.0 -p github.com/hyperledger/fabric/chaincode/ && \
peer chaincode instantiate -C demochannel -n democc -v 1.0 -c '{"Args":["init","a","100","b","200"]}' -o orderer0.${DOMAIN}:5050 --tls  --cafile /var/hyperledger/fabric/tls/ca.crt && \
peer chaincode invoke -C demochannel -n democc -c '{"Args":["invoke","a","b","10"]}' --tls --cafile /var/hyperledger/fabric/tls/ca.crt