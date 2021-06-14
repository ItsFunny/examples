#!/bin/bash


export orderOrgName=demo.com && \
export orderName=orderer.demo.com && \
export orderPort=5050 && \
export ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/${orderOrgName}/orderers/${orderName}/tls/ca.crt && \
export peerIndex=0 && \
export port=11051 && \
export orgIndex=1 && \
export channelName=demochannel1 && \
export CORE_PEER_ADDRESS=peer${peerIndex}.org${orgIndex}.com:${port} && \
export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org${orgIndex}.com/users/Admin@org${orgIndex}.com/msp && \
export CORE_PEER_LOCALMSPID=Org${orgIndex}MSP && \
export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org${orgIndex}.com/peers/peer${peerIndex}.org${orgIndex}.com/tls/ca.crt
 peer chaincode install -n democc -v 1.0 -p github.com/hyperledger/fabric/chaincode/
#peer chaincode instantiate -o ${orderName}:${orderPort} --tls true --cafile ${ORDERER_CA} -C demochannel -n democc -v 1.0 -c '{"Args":["init","a","b","100"]}' -P "OutOf (1,'Org1MSP.peer','Org2MSP.peer','Org3MSP.peer','Org4MSP.peer','Org5MSP.peer')"
# channel 1
limit=1
if [[ -n ${1} ]];then
limit=${1}
fi

peer chaincode query -C demochannel0 -n democc -c '{"Args":["getvalue","a"]}'
peer chaincode invoke -o ${orderName}:${orderPort} --tls true --cafile ${ORDERER_CA} -C demochannel0 -n democc   -c '{"Args":["setvalue","qqqq","b","10"]}'
peer chaincode query -C demochannel0 -n democc -c '{"Args":["getvalue","qqqq"]}'