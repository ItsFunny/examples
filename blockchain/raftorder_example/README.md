# etcd 搭建模式
- 5个order节点


| order节点名称 | 域名 | 端口 |
| :------| ------: | :------: |
| Order0Org | orderer0.demo.com | 7050 |
| UserOrg | peer0.user.vlink.com | 7051 |
| UserOrg | peer1.user.vlink.com | 7061 |


```
1. 创建相关文件
./build.sh
2. 启动docker
cd containers && \
docker-compose -f docker-compose-orderer.yaml  -f docker-compose-cli.yaml up -d
3. 进入cli命令行
docker exec -it cli /bin/bash
4. 执行如下命令行
peer channel create -o orderer.demo.com:7050 -c demochannel -f ./artifacts/demo.tx --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/demo.com/orderers/orderer.demo.com/msp/tlscacerts/tlsca.demo.com-cert.pem && \
peer channel join -b demochannel.block && \
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.demo.com/users/Admin@org2.demo.com/msp CORE_PEER_ADDRESS=peer0.org2.demo.com:9051 CORE_PEER_LOCALMSPID=Org2MSP CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.demo.com/peers/peer0.org2.demo.com/tls/ca.crt peer channel join -b demochannel.block && \
peer channel update -o orderer.demo.com:7050 -c demochannel -f ./artifacts/org1mspanchors.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/demo.com/orderers/orderer.demo.com/msp/tlscacerts/tlsca.demo.com-cert.pem && \
peer chaincode install -n democc -v 1.0 -p github.com/chaincode && \
peer chaincode instantiate -o orderer.demo.com:7050 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/demo.com/orderers/orderer.demo.com/msp/tlscacerts/tlsca.demo.com-cert.pem -C demochannel -n democc -v 1.0 -c '{"Args":["init","a","b","100"]}' -P "AND ('Org1MSP.peer','Org2MSP.peer')" && \
peer chaincode query -C demochannel -n democc -c '{"Args":["query","a"]}'
```

