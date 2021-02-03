#!/usr/bin/env bash

# 第一个参数为 channel 有几个channel,默认为1个demochannel1
# 第二个参数为 org limit,代表有几个org
# 第三个参数为 peerLimit 代表一个org下面有几个peer,默认也是为1个

echo "安装chaincode"
limit=1
if [[  ${2} ]];then
limit=${2}
fi

peerLimit=1
if [[ -n ${3} ]];then
peerLimit=${3}
peerLimit=`expr ${peerLimit} - 1 `
fi

portIndex=10000
for (( i=1; i<=${limit}; i++ ))
do
    orgUp=`expr 1000 \* ${i} `
    orgUp=`expr ${orgUp} + 51 `
    portIndex=`expr ${portIndex} + ${orgUp} `
    for (( j=0; j<=${peerLimit}; j++ ))
    do
        peerIndex=${j}
        up=`expr 10 \* ${j}`
        portIndex=`expr ${portIndex} + ${up} `
        echo -e ">>>>>>>>>>>>>>>>>>peer:peer${peerIndex}.org${i}.com:${portIndex} 开始安装chaincode>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>"
        export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org${i}.com/users/Admin@org${i}.com/msp  CORE_PEER_ADDRESS=peer${peerIndex}.org${i}.com:${portIndex} CORE_PEER_LOCALMSPID=Org${i}MSP CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org${i}.com/peers/peer${peerIndex}.org${i}.com/tls/ca.crt
	for file in `ls ${GOPATH}/src/bidchain/chaincode/`
	do
   	   echo "peer chaincode install -n democc -v 1.0 -p bidchain/chaincode/${file}"
	   peer chaincode install -n democc -v 1.0 -p bidchain/chaincode/${file}
	done
        echo -e  "<<<<<<<<<<<<<<<<<peer:peer${peerIndex}.org${i}.com:${portIndex} 结束安装chaincode<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<\n\n"
    done
    portIndex=10000
done

./scripts/instantitatechaincode.sh ${1}
