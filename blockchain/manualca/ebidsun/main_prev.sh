#!/usr/bin/env bash


export DOMAIN=bidsun.com && \
peer channel create -o orderer.${DOMAIN}:7050 -c demo-channel -f ./artifacts/channel.tx --tls true --cafile /var/hyperledger/fabric/orderers/orderer.bidsun.com/tls/ca.crt && \
sleep 10 && \
peer channel join -b ${PWD}/demo-channel.block && \
peer chaincode install -n democc -v 1.0 -p github.com/hyperledger/fabric/chaincode/ && \
peer chaincode instantiate -C demo-channel -n democc -v 1.0 -c '{"Args":["init","a","100","b","200"]}' -o orderer.${DOMAIN}:7050 --tls --cafile /var/hyperledger/fabric/orderers/orderer.bidsun.com/msp/tlscacerts/tlsca.bidsun.com-cert.pem && \
peer chaincode invoke -C demochannel -n democc -c '{"Args":["invoke","a","b","10"]}' --tls --cafile /var/hyperledger/fabric/tls/ca.crt