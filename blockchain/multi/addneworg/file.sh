#!/bin/bash


# 1. 生成配置信息
cryptogen_multi generate --config=crypto-config.yaml

# 2. 使用configtxgen 生成json文件
export FABRIC_CFG_PATH=${pwd} && configtxgen_multi -printOrg Org6MSP > ./org6.json
# 3. 将这个peerOrganizations cp 到上级的crypto-config 里 并且将这个json文件挪到 artifacts中
cp -r ./crypto-config/peerOrganizations/org6.com ../../crypto-config/peerOrganizations/
cp ./org6.json ../../artifacts/


# 4. 准备下cli 环境
docker exec -it cli /bin/bash
# 5. 设置基本环境:
export ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/demo.com/orderers/orderer.demo.com/tls/ca.crt
export CHANNEL_NAME=demochannel1
# 6. 拉取到最新的配置块
peer channel fetch config config_block.pb -o orderer.demo.com:5050 -c $CHANNEL_NAME --tls --cafile $ORDERER_CA
# 7. 将proto 文件转换位 json文件
configtxlator proto_decode --input config_block.pb --type common.Block | jq .data.data[0].payload.data.config > config.json
# 8. 利用jq 将数据写入到配置区块的json文件中
jq -s '.[0] * {"channel_group":{"groups":{"Application":{"groups": {"Org6MSP":.[1]}}}}}' config.json ./artifacts/org6.json > modified_config.json
# 9. 将已经修改过的json文件转为proto 文件
configtxlator proto_encode --input modified_config.json --type common.Config --output modified_config.pb
# 10. configtxlator计算前后的差值,并且生成一个新的文件
configtxlator compute_update --channel_id $CHANNEL_NAME --original config.pb --updated modified_config.pb --output org6_update.pb
# 11. submit 之前还需要再转为json
configtxlator proto_decode --input org6_update.pb --type common.ConfigUpdate | jq . > org6_update.json
# 12. 将数据包装成数字信封的格式
echo '{"payload":{"header":{"channel_header":{"channel_id":"'$CHANNEL_NAME'", "type":2}},"data":{"config_update":'$(cat org6_update.json)'}}}' | jq . > org6_update_in_envelope.json
# 13. 最后再转为protobuf格式,准备submit 提案
configtxlator proto_encode --input org6_update_in_envelope.json --type common.Envelope --output org6_update_in_envelope.pb
# 14. 对protobuf 的数字信封进行签名
peer channel signconfigtx -f org6_update_in_envelope.pb
# 15. 或许还需要转换为其他组织 进行签名?

# 16 发送channel update的消息
peer channel update -f org6_update_in_envelope.pb -c $CHANNEL_NAME -o orderer.demo.com:5050 --tls --cafile $ORDERER_CA

# 遇到的问题
configtxgen 需要FABRIC_CFG_PATH 的原因:?