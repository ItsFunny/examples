#!/usr/bin/env bash

# 添加一个组织到channel中
## 在这之前,先搭建好环境,既,将create.sh 中的执行一遍,除了测试,再将该文件执行一遍
## 在启动org3的cli
## 不足的地方,加入channel 不应该是client 操作,而应该是server端操作,所以 cli的target peer不应该是新加入的节点
## 而是原先服务器的节点才是

# 1. 需要创建一个org3 的crypt-config.yaml 和configtx.yaml
cd org3/
rm -rf artifacts
rm -rf crypto-config
cryptogen generate --config=./org3-crypto.yaml
# 2. 生成msp等
export FABRIC_CFG_PATH=${pwd}

configtxgen -printOrg Org3MSP > org3.json

mkdir artifacts

cd ../ && cp -r ./crypto-config/ordererOrganizations ./org3/crypto-config/

# 将需要签名的组织的文件复制到新的组织下
cp -r ./crypto-config/peerOrganizations/* ./org3/crypto-config/peerOrganizations/
mv ./org3/crypto-config ./org3/artifacts/

# 将新的组织的json文件移动到统一文件下
mv ./org3/org3.json ./org3/artifacts/crypto-config/


cd ./org3/

docker-compose -f docker-compose-org3-cli.yaml --project-name containers up -d

#cd org3/ && rm -rf artifacts  && cryptogen generate --config=./org3-crypto.yaml && export FABRIC_CFG_PATH=${pwd} && configtxgen -printOrg Org3MSP > org3.json && mkdir artifacts && mv org3.json ./artifacts && cd ../ && cp -r ./crypto-config/ordererOrganizations ./org3/crypto-config/ && cp -r ./crypto-config/peerOrganizations/* ./org3/crypto-config/peerOrganizations/ && mv ./org3/crypto-config ./org3/artifacts/


# 进入cli 容器
docker exec -it Org3cli /bin/bash
# 更新并且安装jq
apt-get -y update && apt-get -y install jq


# 临时设置 order 的全局变量 并且 拉取channel最新的配置块  因为需要拉取的是order的信息,所以需要配置order的msp等信息,而且需要为临时设置变量
CORE_PEER_LOCALMSPID=OrdererMSP CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/demo.com/orderers/orderer.demo.com/msp/tlscacerts/tlsca.demo.com-cert.pem CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/demo.com/users/Admin@demo.com/msp peer channel fetch config config_block.pb -o orderer.demo.com:7050 -c demochannel --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/demo.com/orderers/orderer.demo.com/msp/tlscacerts/tlsca.demo.com-cert.pem

# 解码为json格式
configtxlator proto_decode --input config_block.pb --type common.Block | jq .data.data[0].payload.data.config > config.json

# 将新的组织的json文件与之前的json文件合并
jq -s '.[0] * {"channel_group":{"groups":{"Application":{"groups": {"Org3MSP":.[1]}}}}}' config.json ./crypto/org3.json >modified_config.json

# 将新的json文件转成二进制pb 文件
 configtxlator proto_encode --input modified_config.json --type common.Config --output modified_config.pb
 # 将之前的配置json文件转换成二进制pb文件
configtxlator proto_encode --input config.json --type common.Config --output config.pb

# 合并2个 新旧配置的区别,并合并生成新的配置文件
configtxlator compute_update --channel_id demochannel --original config.pb --updated modified_config.pb --output org3_update.pb

# 将新生成的文件转成json 可读
configtxlator proto_decode --input org3_update.pb --type common.ConfigUpdate | jq . > org3_update.json

# 通过新的json文件,将该文件转成envelop 对象
echo '{"payload":{"header":{"channel_header":{"channel_id":"demochannel", "type":2}},"data":{"config_update":'$(cat org3_update.json)'}}}' | jq . > org3_update_in_envelope.json
# 将这个信封文件转成protobuf
configtxlator proto_encode --input org3_update_in_envelope.json --type common.Envelope --output org3_update_in_envelope.pb


# 准备发起请求,发起请求前需要相关的组织进行签名
peer channel signconfigtx -f org3_update_in_envelope.pb
# 第1个组织签名
CORE_PEER_LOCALMSPID=Org1MSP CORE_PEER_ADDRESS=peer0.org1.demo.com:8051 CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.demo.com/peers/peer0.org1.demo.com/tls/ca.crt CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.demo.com/users/Admin@org1.demo.com/msp peer channel signconfigtx -f org3_update_in_envelope.pb
# 第2个组织签名
CORE_PEER_LOCALMSPID=Org2MSP CORE_PEER_ADDRESS=peer0.org2.demo.com:9051 CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.demo.com/peers/peer0.org2.demo.com/tls/ca.crt CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.demo.com/users/Admin@org2.demo.com/msp peer channel signconfigtx -f org3_update_in_envelope.pb
# 发起请求,注意发起请求的时候也需要传递 msp等信息,并且是order的配置
CORE_PEER_LOCALMSPID=OrdererMSP CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/demo.com/orderers/orderer.demo.com/msp/tlscacerts/tlsca.demo.com-cert.pem CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/demo.com/users/Admin@demo.com/msp peer channel update -f org3_update_in_envelope.pb -c demochannel -o orderer.demo.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/demo.com/users/Admin@demo.com/msp peer channel fetch config config_block.pb -o orderer.demo.com:7050 -c demochannel --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/demo.com/orderers/orderer.demo.com/msp/tlscacerts/tlsca.demo.com-cert.pem

# 准备加入该channel: 获取创世快
peer channel fetch 0 demo.block -o orderer.demo.com:7050 -c demochannel --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/demo.com/orderers/orderer.demo.com/msp/tlscacerts/tlsca.demo.com-cert.pem
# 加入channel
 peer channel join -b demo.block

 # 大功告成

# 测试安装合约
peer chaincode install -n democc -v 2.0 -p github.com/chaincode/

## 因为是个人,所以直接切换到其他已经存在的组织,不然是无法升级组织的
## 组织1 安装版本2.0
CORE_PEER_LOCALMSPID=Org1MSP CORE_PEER_ADDRESS=peer0.org1.demo.com:8051 CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.demo.com/peers/peer0.org1.demo.com/tls/ca.crt CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.demo.com/users/Admin@org1.demo.com/msp peer chaincode install -n democc -v 2.0 -p github.com/chaincode/
## 组织2 安装版本2.0
CORE_PEER_LOCALMSPID=Org2MSP CORE_PEER_ADDRESS=peer0.org2.demo.com:9051 CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.demo.com/peers/peer0.org2.demo.com/tls/ca.crt CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.demo.com/users/Admin@org2.demo.com/msp peer chaincode install -n democc -v 2.0 -p github.com/chaincode/
## 组织1 升级版本成2.0
# 因为版本为2.0 ,同时新增了一个组织,policy 更改为也需要该组织背书,则需要升级(policy就算不变也需要升级,不然低版本没安装是无法调用的)
CORE_PEER_LOCALMSPID=Org1MSP CORE_PEER_ADDRESS=peer0.org1.demo.com:8051 CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.demo.com/peers/peer0.org1.demo.com/tls/ca.crt CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.demo.com/users/Admin@org1.demo.com/msp peer chaincode upgrade -o orderer.demo.com:7050 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/demo.com/orderers/orderer.demo.com/msp/tlscacerts/tlsca.demo.com-cert.pem -C demochannel -n democc -v 2.0 -c '{"Args":["init","a","90","b","210"]}' -P "OR ('Org1MSP.peer','Org2MSP.peer','Org3MSP.peer')"
# 测试调用合约 (不用单独实例化,直接调用即可,会自动实例化一个)
peer chaincode query -C demochannel -n democc -c '{"Args":["query","a"]}'

## 注意,如果安装的时候版本号不一致 ,既原先的democc 升级了,此时如果query,低版本原先的未安装的话,还是无法调用的

# anchor peer 准备加入到网络通信中,这样下次通信都能发现这个新的组织的anchor peer
peer channel fetch config config_block.pb -o orderer.demo.com:7050 -c demochannel --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/demo.com/orderers/orderer.demo.com/msp/tlscacerts/tlsca.demo.com-cert.pem
# 转成json 化
configtxlator proto_decode --input config_block.pb --type common.Block | jq .data.data[0].payload.data.config > config.json
# 通过jq 加入到数据中
jq '.channel_group.groups.Application.groups.Org3MSP.values += {"AnchorPeers":{"mod_policy": "Admins","value":{"anchor_peers": [{"host": "peer0.org3.demo.com","port": 21051}]},"version": "0"}}' config.json > modified_anchor_config.json
# 将之前的配置文件转成通信的pb
configtxlator proto_encode --input config.json --type common.Config --output config.pb
# 将新的配置json文件再转为pb,用于计算合并差值
configtxlator proto_encode --input modified_anchor_config.json --type common.Config --output modified_anchor_config.pb
# 计算合并差值,生成新的文件
configtxlator compute_update --channel_id demochannel --original config.pb --updated modified_anchor_config.pb --output anchor_update.pb

# 新的差值文件,转成json 准备封装成envelop
configtxlator proto_decode --input anchor_update.pb --type common.ConfigUpdate | jq . > anchor_update.json
# 封装成envelop
echo '{"payload":{"header":{"channel_header":{"channel_id":"demochannel", "type":2}},"data":{"config_update":'$(cat anchor_update.json)'}}}' | jq . > anchor_update_in_envelope.json
# envelop转成pb对象用于通讯
configtxlator proto_encode --input anchor_update_in_envelope.json --type common.Envelope --output anchor_update_in_envelope.pb
# 发送更新请求
peer channel update -f anchor_update_in_envelope.pb -c demochannel -o orderer.demo.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/demo.com/orderers/orderer.demo.com/msp/tlscacerts/tlsca.demo.com-cert.pem





-------
rm -rf crypto-config && \
cryptogen generate --config=./org3-crypto.yaml && \
export FABRIC_CFG_PATH=${PWD} && \
configtxgen -printOrg Org3MSP > org3.json && \
mkdir artifacts && cd ../ && \
cp -r ./crypto-config/ordererOrganizations ./org3/crypto-config/ && \
cp -r ./crypto-config/peerOrganizations/* ./org3/crypto-config/peerOrganizations/ && \
mv ./org3/crypto-config ./org3/artifacts/ && \
mv ./org3/org3.json ./org3/artifacts/crypto-config/ && \
cd ./org3/ && \
docker-compose -f docker-compose-org3-cli.yaml --project-name containers up -d && \
docker exec -it Org3cli /bin/bash

# 上面的汇总命令行
apt-get -y update && apt-get -y install jq && \
CORE_PEER_LOCALMSPID=OrdererMSP CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/demo.com/orderers/orderer.demo.com/msp/tlscacerts/tlsca.demo.com-cert.pem CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/demo.com/users/Admin@demo.com/msp peer channel fetch config config_block.pb -o orderer.demo.com:7050 -c demochannel --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/demo.com/orderers/orderer.demo.com/msp/tlscacerts/tlsca.demo.com-cert.pem && \
configtxlator proto_decode --input config_block.pb --type common.Block | jq .data.data[0].payload.data.config > config.json && \
jq -s '.[0] * {"channel_group":{"groups":{"Application":{"groups": {"Org3MSP":.[1]}}}}}' config.json ./crypto/org3.json >modified_config.json && \
configtxlator proto_encode --input modified_config.json --type common.Config --output modified_config.pb && \
configtxlator proto_encode --input config.json --type common.Config --output config.pb && \
configtxlator compute_update --channel_id demochannel --original config.pb --updated modified_config.pb --output org3_update.pb && \
configtxlator proto_decode --input org3_update.pb --type common.ConfigUpdate | jq . > org3_update.json && \
echo '{"payload":{"header":{"channel_header":{"channel_id":"demochannel", "type":2}},"data":{"config_update":'$(cat org3_update.json)'}}}' | jq . > org3_update_in_envelope.json && \
configtxlator proto_encode --input org3_update_in_envelope.json --type common.Envelope --output org3_update_in_envelope.pb && \
peer channel signconfigtx -f org3_update_in_envelope.pb && \
CORE_PEER_LOCALMSPID=Org1MSP CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.demo.com/peers/peer0.org1.demo.com/tls/ca.crt CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.demo.com/users/Admin@org1.demo.com/msp peer channel signconfigtx -f org3_update_in_envelope.pb && \
CORE_PEER_LOCALMSPID=Org2MSP CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.demo.com/peers/peer0.org2.demo.com/tls/ca.crt CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.demo.com/users/Admin@org2.demo.com/msp peer channel signconfigtx -f org3_update_in_envelope.pb && \
CORE_PEER_LOCALMSPID=OrdererMSP CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/demo.com/orderers/orderer.demo.com/msp/tlscacerts/tlsca.demo.com-cert.pem CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/demo.com/users/Admin@demo.com/msp peer channel update -f org3_update_in_envelope.pb -c demochannel -o orderer.demo.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/demo.com/users/Admin@demo.com/msp peer channel fetch config config_block.pb -o orderer.demo.com:7050 -c demochannel --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/demo.com/orderers/orderer.demo.com/msp/tlscacerts/tlsca.demo.com-cert.pem && \
peer channel fetch 0 demo.block -o orderer.demo.com:7050 -c demochannel --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/demo.com/orderers/orderer.demo.com/msp/tlscacerts/tlsca.demo.com-cert.pem && \
peer channel join -b demo.block && \
echo "新增组织成功" && \
peer channel fetch 0 demo.block -o orderer.demo.com:7050 -c demochannel --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/demo.com/orderers/orderer.demo.com/msp/tlscacerts/tlsca.demo.com-cert.pem && \
peer channel join -b demo.block && \
peer chaincode install -n democc -v 2.0 -p github.com/chaincode/ && \
CORE_PEER_LOCALMSPID=Org1MSP CORE_PEER_ADDRESS=peer0.org1.demo.com:8051 CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.demo.com/peers/peer0.org1.demo.com/tls/ca.crt CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.demo.com/users/Admin@org1.demo.com/msp peer chaincode install -n democc -v 2.0 -p github.com/chaincode/ && \
CORE_PEER_LOCALMSPID=Org2MSP CORE_PEER_ADDRESS=peer0.org2.demo.com:9051 CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.demo.com/peers/peer0.org2.demo.com/tls/ca.crt CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.demo.com/users/Admin@org2.demo.com/msp peer chaincode install -n democc -v 2.0 -p github.com/chaincode/ && \
CORE_PEER_LOCALMSPID=Org1MSP CORE_PEER_ADDRESS=peer0.org1.demo.com:8051 CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.demo.com/peers/peer0.org1.demo.com/tls/ca.crt CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.demo.com/users/Admin@org1.demo.com/msp peer chaincode upgrade -o orderer.demo.com:7050 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/demo.com/orderers/orderer.demo.com/msp/tlscacerts/tlsca.demo.com-cert.pem -C demochannel -n democc -v 2.0 -c '{"Args":["init","a","90","b","210"]}' -P "OR ('Org1MSP.peer','Org2MSP.peer','Org3MSP.peer')" && \
peer chaincode query -C demochannel -n democc -c '{"Args":["query","a"]}' && \
peer channel fetch config config_block.pb -o orderer.demo.com:7050 -c demochannel --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/demo.com/orderers/orderer.demo.com/msp/tlscacerts/tlsca.demo.com-cert.pem && \
echo '转成json 化' && \
configtxlator proto_decode --input config_block.pb --type common.Block | jq .data.data[0].payload.data.config > config.json && \
echo '通过jq 加入到数据中' && \
jq '.channel_group.groups.Application.groups.Org3MSP.values += {"AnchorPeers":{"mod_policy": "Admins","value":{"anchor_peers": [{"host": "peer0.org3.demo.com","port": 21051}]},"version": "0"}}' config.json > modified_anchor_config.json && \
echo  '将之前的配置文件转成通信的pb' && \
configtxlator proto_encode --input config.json --type common.Config --output config.pb && \
echo '将新的配置json文件再转为pb,用于计算合并差值' && \
configtxlator proto_encode --input modified_anchor_config.json --type common.Config --output modified_anchor_config.pb && \
echo  '计算合并差值,生成新的文件' && \
configtxlator compute_update --channel_id demochannel --original config.pb --updated modified_anchor_config.pb --output anchor_update.pb && \
echo '新的差值文件,转成json 准备封装成envelop' && \
configtxlator proto_decode --input anchor_update.pb --type common.ConfigUpdate | jq . > anchor_update.json && \
echo '封装成envelop' && \
echo '{"payload":{"header":{"channel_header":{"channel_id":"demochannel", "type":2}},"data":{"config_update":'$(cat anchor_update.json)'}}}' | jq . > anchor_update_in_envelope.json && \
echo 'envelop转成pb对象用于通讯' && \
configtxlator proto_encode --input anchor_update_in_envelope.json --type common.Envelope --output anchor_update_in_envelope.pb && \
echo '发送更新请求' && \
peer channel update -f anchor_update_in_envelope.pb -c demochannel -o orderer.demo.com:7050 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/demo.com/orderers/orderer.demo.com/msp/tlscacerts/tlsca.demo.com-cert.pem
