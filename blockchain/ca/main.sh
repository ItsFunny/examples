#!/usr/bin/env bash


export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/admin/msp

peer channel create -o orderer0.demo.com:7050 -c demochannel -f ../artifacts/demochannel.tx --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/tls-msp/tlscacerts/tls-0-0-0-0-7052.pem  && \