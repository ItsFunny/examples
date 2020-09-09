#!/usr/bin/env bash


export DOMAIN=bidsun.com && \
peer channel create -o orderer.${DOMAIN}:7050 -c demo-channel -f ./artifacts/channel.tx --tls true --cafile /var/hyperledger/fabric/orderers/orderer.bidsun.com/tls/ca.crt && \
sleep 10 && \
peer channel join -b ${PWD}/demo-channel.block && \

peer lifecycle chaincode package demo2.tar.gz  --path github.com/hyperledger/fabric/chaincode/ --label demo_2.0
peer lifecycle chaincode install demo2.tar.gz
peer lifecycle chaincode queryinstalled


# 只有approve 之后才可以commit
peer lifecycle chaincode approveformyorg -o orderer.bidsun.com:7050 --channelID demo-channel --signature-policy "OR('BidsunMSP.admin', 'BidsunMSP.peer', 'BidsunMSP.client')"  --name democc --version 1 --init-required --sequence 1 --package-id demo_2.0 --waitForEvent --tls --cafile /var/hyperledger/fabric/orderers/orderer.bidsun.com/msp/tlscacerts/tlsca.bidsun.com-cert.pem


peer lifecycle chaincode commit -o orderer.bidsun.com:7050 --channelID demo-channel --name democc --version 1.0 --sequence 1 --init-required --tls --cafile /var/hyperledger/fabric/orderers/orderer.bidsun.com/msp/tlscacerts/tlsca.bidsun.com-cert.pem --peerAddresses peer0.bidsun.com:7051

peer chaincode instantiate -C demochannel -n democc -v 1.0 -c '{"Args":["init","a","100","b","200"]}' -o orderer0.${DOMAIN}:5050 --tls  --cafile /var/hyperledger/fabric/tls/ca.crt && \
peer chaincode invoke -C demochannel -n democc -c '{"Args":["invoke","a","b","10"]}' --tls --cafile /var/hyperledger/fabric/tls/ca.crt