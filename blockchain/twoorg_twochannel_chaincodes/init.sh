#!/bin/bash

rm -rf vendor
govendor init
govendor add +e

cp -r \
   "${GOPATH}/src/github.com/ethereum/go-ethereum/crypto/secp256k1/libsecp256k1" \
   "${GOPATH}/src/examples/blockchain/twoorg_twochannel_chaincodes/vendor/github.com/ethereum/go-ethereum/crypto/secp256k1/"