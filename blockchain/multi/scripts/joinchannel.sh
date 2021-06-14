#!/usr/bin/env bash

# 第一个参数为 channel 有几个channel,默认为1个demochannel1
# 第二个参数为 org limit,代表有几个org
# 第三个参数为 peerLimit 代表一个org下面有几个peer,默认也是为1个
cd ${PWD}

channelLimit=1
if [[ -n ${1} ]];then
    channelLimit=${1}
fi

limit=1
if [[  ${2} ]];then
limit=${2}
fi

peerLimit=0
if [[ -n ${3} ]];then
peerLimit=${3}
fi

sleep 5
export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.com/users/Admin@org1.com/msp
export ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/demo.com/orderers/orderer.demo.com/tls/ca.crt


for (( i=0; i<${channelLimit}; i++))
do
    peer channel create -o orderer.demo.com:5050 -c demochannel${i} -f ../peer/artifacts/demochannel${i}.tx --tls true --cafile ${ORDERER_CA}
    sleep 10
done



for (( i=0; i<${channelLimit}; i++ ))
do
    portIndex=10000
    for (( j=0; j<${limit}; j++ ))
    do
        orgUp=`expr 1000 \* ${j} `
        orgUp=`expr ${orgUp} + 51 `
        portIndex=`expr ${portIndex} + ${orgUp} `
        for (( k=0; k<${peerLimit}; k++ ))
        do
            peerIndex=${k}
            up=`expr 10 \* ${k}`
            portIndex=`expr ${portIndex} + ${up} `
            echo "组织 [Org${j}MSP,peer${peerIndex}.org${j}.com:${portIndex}]加入到demochannel${i}中"
            export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org${j}.com/users/Admin@org${j}.com/msp  CORE_PEER_ADDRESS=peer${peerIndex}.org${j}.com:${portIndex} CORE_PEER_LOCALMSPID=Org${j}MSP CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org${j}.com/peers/peer${peerIndex}.org${j}.com/tls/ca.crt
            peer channel join -b ${PWD}/demochannel${i}.block
            if [[ $? -ne 0 ]]; then
                sleep 3
                echo ">>>>>>><<<<<<<<<加入失败<<<<<>>>>>>>>>>>>><<<<<"
                exit -1
            fi
        done
        portIndex=10000
    done
done

for (( i=0; i<${channelLimit}; i++ ))
do
    for (( j=0; j<${limit}; j++ ))
    do
        orgUp=`expr 1000 \* ${j} `
        orgUp=`expr ${orgUp} + 51 `

         echo "组织 [Org${j}MSP 指定anchorpeer中"
        CHANNEL_NAME=demochannel${i}
        export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org${j}.com/users/Admin@org${j}.com/msp  CORE_PEER_ADDRESS=peer0.org${j}.com:1${j}051 CORE_PEER_LOCALMSPID=Org${j}MSP CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org${j}.com/peers/peer0.org${j}.com/tls/ca.crt
        peer channel update -o orderer.demo.com:5050 -c ${CHANNEL_NAME} -f ../peer/artifacts/org${j}mspanchors.tx --tls --cafile ${ORDERER_CA}
    done
done


./scripts/installchaincode.sh ${1} ${2} ${3}
