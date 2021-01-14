#!/bin/bash


# 第一个参数为需要几个org 签名
# 第2个参数为channel名称
orgLimit=5
if [[ -n ${1} ]];then
orgLimit=${1}
fi
channelName=demochannel1
if [[ -n ${2} ]];then
channelName=${2}
fi
echo "orgLimit=${orgLimit},channelName=${channelName}"
#
# 5. 设置基本环境:
export ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/demo.com/orderers/orderer.demo.com/tls/ca.crt
export CHANNEL_NAME=${channelName}
# 6. 拉取到最新的配置块
peer channel fetch config config_block.pb -o orderer.demo.com:5050 -c $CHANNEL_NAME --tls --cafile $ORDERER_CA
# 7. 将proto 文件转换位 json文件
configtxlator proto_decode --input config_block.pb --type common.Block | jq .data.data[0].payload.data.config > config.json
# 8. 利用jq 将数据写入到配置区块的json文件中
jq -s '.[0] * {"channel_group":{"groups":{"Application":{"groups": {"Org6MSP":.[1]}}}}}' config.json ./artifacts/org6.json > modified_config.json
# 9. 将已经修改过的json文件转为proto 文件
configtxlator proto_encode --input modified_config.json --type common.Config --output modified_config.pb
# 10. configtxlator计算前后的差值,并且生成一个新的文件
configtxlator compute_update --channel_id $CHANNEL_NAME --original config_block.pb --updated modified_config.pb --output org6_update.pb
# 11. submit 之前还需要再转为json
configtxlator proto_decode --input org6_update.pb --type common.ConfigUpdate | jq . > org6_update.json
# 12. 将数据包装成数字信封的格式
echo '{"payload":{"header":{"channel_header":{"channel_id":"'$CHANNEL_NAME'", "type":2}},"data":{"config_update":'$(cat org6_update.json)'}}}' | jq . > org6_update_in_envelope.json
# 13. 最后再转为protobuf格式,准备submit 提案
configtxlator proto_encode --input org6_update_in_envelope.json --type common.Envelope --output org6_update_in_envelope.pb
# 14,15. 对protobuf 的数字信封进行签名 或许还需要转换为其他组织 进行签名?
portIndex=10000
for (( j=1; j<=${orgLimit}; j++ ))
do
    orgUp=`expr 1000 \* ${j} `
    orgUp=`expr ${orgUp} + 51 `
    portIndex=`expr ${portIndex} + ${orgUp} `
    for (( k=0; k<=0; k++ ))
    do
        peerIndex=${k}
        up=`expr 10 \* ${k}`
        portIndex=`expr ${portIndex} + ${up} `
        export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org${j}.com/users/Admin@org${j}.com/msp  CORE_PEER_ADDRESS=peer${peerIndex}.org${j}.com:${portIndex} CORE_PEER_LOCALMSPID=Org${j}MSP CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org${j}.com/peers/peer${peerIndex}.org${j}.com/tls/ca.crt
        echo " export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org${j}.com/users/Admin@org${j}.com/msp  CORE_PEER_ADDRESS=peer${peerIndex}.org${j}.com:${portIndex} CORE_PEER_LOCALMSPID=Org${j}MSP CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org${j}.com/peers/peer${peerIndex}.org${j}.com/tls/ca.crt"
        peer channel signconfigtx -f org6_update_in_envelope.pb
    done
    portIndex=10000
done

# 16 发送channel update的消息
peer channel update -f org6_update_in_envelope.pb -c $CHANNEL_NAME -o orderer.demo.com:5050 --tls --cafile $ORDERER_CA
#