#!/usr/bin/env bash


export ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/demo.com/orderers/orderer.demo.com/tls/ca.crt


for i in {1..5}
do
    echo "组织${i}实例化chaincode"
    export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org${i}.com/users/Admin@org${i}.com/msp  CORE_PEER_ADDRESS=peer0.org${i}.com:1${i}051 CORE_PEER_LOCALMSPID=Org${i}MSP CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org${i}.com/peers/peer0.org${i}.com/tls/ca.crt
    echo ${CORE_PEER_MSPCONFIGPATH}
    echo ${CORE_PEER_ADDRESS}
    echo ${CORE_PEER_LOCALMSPID}
    echo ${CORE_PEER_TLS_ROOTCERT_FILE}
     peer chaincode instantiate -o orderer.demo.com:5050 --tls true --cafile ${ORDERER_CA} -C demochannel -n democc -v 1.0 -c '{"Args":["init","a","b","100"]}' -P "OutOf (2,'Org1MSP.peer','Org2MSP.peer','Org3MSP.peer','Org4MSP.peer','Org5MSP.peer')"
#    peer chaincode instantiate -o orderer.demo.com:5050 --tls true --cafile ${ORDERER_CA} -C demochannel -n democc -v 1.0 -c '{"Args":["init","a","b","100"]}' -P "AND ('Org1MSP.peer','Org2MSP.peer','Org3MSP.peer','Org4MSP.peer','Org5MSP.peer')"
done

#peer chaincode invoke -C demochannel -n democc -c '{"Args":["invoke","a","b","10"]}' --tls --cafile ${ORDERER_CA}