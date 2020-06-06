#!/usr/bin/env bash


peer channel create -o ord-erer0.demo.com:7050 -c demochannel -f ./artifacts/demo.tx --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/demo.com/orderers/orderer0.demo.com/msp/tlscacerts/tlsca.demo.com-cert.pem