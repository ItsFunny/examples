#!/bin/bash
set -eu

channelName=$1
ORDERER_ADDRESS=$2
ORDERER_CA=$3
newOrdererHost=$4
newOrdererPort=$5

setOrdererGlobals() {
   CORE_PEER_ADDRESS=orderer.bidsun.com:7050
   CORE_PEER_LOCALMSPID=OrdererMSP
   CORE_PEER_TLS_ENABLED=true
   CORE_PEER_TLS_CERT_FILE=/opt/workspace/crypto/ordererOrganizations/bidsun.com/orderers/orderer.bidsun.com/tls/server.crt
   CORE_PEER_TLS_KEY_FILE=/opt/workspace/crypto/ordererOrganizations/bidsun.com/orderers/orderer.bidsun.com/tls/server.crt 
   CORE_PEER_TLS_ROOTCERT_FILE=/opt/workspace/crypto/ordererOrganizations/bidsun.com/orderers/orderer.bidsun.com/tls/ca.crt
   CORE_PEER_MSPCONFIGPATH=/opt/workspace/crypto/ordererOrganizations/bidsun.com/users/Admin@bidsun.com/msp/
   ORDERER_CA=/opt/workspace/crypto/ordererOrganizations/bidsun.com/orderers/orderer.bidsun.com/msp/tlscacerts/tlsca.bidsun.com-cert.pem
}

signNewOrderer() {
  configtxlator proto_encode --input config.json --type common.Config --output config.pb
  configtxlator proto_encode --input modified_config.json --type common.Config --output modified_config.pb


  NEW_ORDERER=orderer
  configtxlator compute_update --channel_id $CHANNEL_NAME --original config.pb --updated modified_config.pb --output $NEW_ORDERER"_update.pb"
  configtxlator proto_decode --input $NEW_ORDERER"_update.pb" --type common.ConfigUpdate | jq . > $NEW_ORDERER"_update.json"

  echo '{"payload":{"header":{"channel_header":{"channel_id":"'$CHANNEL_NAME'", "type":2}},"data":{"config_update":'$(cat $NEW_ORDERER"_update.json")'}}}' | jq . > $NEW_ORDERER"_update_in_envelope.json"

  configtxlator proto_encode --input $NEW_ORDERER"_update_in_envelope.json" --type common.Envelope --output $NEW_ORDERER"_update_in_envelope.pb"
  peer channel update -f $NEW_ORDERER"_update_in_envelope.pb" -c $CHANNEL_NAME -o $ORDERER_ADDRESS  --tls --cafile $ORDERER_CA
}

echo "================fetch system channel config==================="
setOrdererGlobals

base64Flag="-w 0"
echo "============= generate orderer.json=============================="
echo "{\"client_tls_cert\":\"$(cat $newOrdererTlsCertFile | base64 $base64Flag)\",\"host\":\"$newOrdererHost\",\"port\":$newOrdererPort,\"server_tls_cert\":\"$(cat $newOrdererTlsCertFile | base64 $base64Flag)\"}" > orderer.json


removeOrdererFromSystemChannel() {
    CHANNEL_NAME=bidchain-sys-channel 
    peer channel fetch config config_block.pb -o $ORDERER_ADDRESS -c $CHANNEL_NAME --tls --cafile $ORDERER_CA
    configtxlator proto_decode --input config_block.pb --type common.Block | jq .data.data[0].payload.data.config > config.json

    echo "=================create config_update.json===================="
    host=$newOrdererHost
    port=$newOrdererPort
    index=$(jq --arg host $host --arg port $port '[ .channel_group.groups.Orderer.values.ConsensusType.value.metadata.consenters  | .[] | .host==$host and .port == ($port|tonumber) ] | index(true)' config.json )
    jq "del(.channel_group.groups.Orderer.values.ConsensusType.value.metadata.consenters[$index])" config.json > tmp.json 
    # 删除指定位置orderer节点
    hp="${host}:${port}"
    index=$(jq --arg host $host --arg port $port '[ .channel_group.values.OrdererAddresses.value.addresses  | .[] | . == "'$hp'" ] | index(true)' tmp.json )
    jq "del(.channel_group.values.OrdererAddresses.value.addresses[$index])"  tmp.json  > modified_config.json
    echo "==============system channel $CHANNEL_NAME remove orderer cert and orderer address======================="
    signNewOrderer

    echo "============= fetch latest system channel config========================"
    peer channel fetch config config_block.pb  -o $ORDERER_ADDRESS  -c $CHANNEL_NAME --tls --cafile $ORDERER_CA
    cp config_block.pb channel-artifacts/$newOrdererHost"_system-channel.block" 

}

removeOrdererFromApplicationChannel() {
    CHANNEL_NAME=$1
    peer channel fetch config config_block.pb -o $ORDERER_ADDRESS -c $CHANNEL_NAME --tls --cafile $ORDERER_CA
    configtxlator proto_decode --input config_block.pb --type common.Block | jq .data.data[0].payload.data.config > config.json

    echo "=================create config_update.json===================="
    host=$newOrdererHost
    port=$newOrdererPort
    index=$(jq --arg host $host --arg port $port '[ .channel_group.groups.Orderer.values.ConsensusType.value.metadata.consenters  | .[] | .host==$host and .port == ($port|tonumber) ] | index(true)' config.json )
    jq "del(.channel_group.groups.Orderer.values.ConsensusType.value.metadata.consenters[$index])" config.json > tmp.json 
    #jq "del(.channel_group.groups.Orderer.values.ConsensusType.value.metadata.consenters[$index])" config.json > modified_config.json 
    # 删除指定位置orderer节点
    hp="${host}:${port}"
    index=$(jq --arg host $host --arg port $port '[ .channel_group.values.OrdererAddresses.value.addresses  | .[] | . == "'$hp'" ] | index(true)' tmp.json )
    jq "del(.channel_group.values.OrdererAddresses.value.addresses[$index])"  tmp.json  > modified_config.json

    #index=$(jq --arg host $host --arg port $port '[ .channel_group.values.OrdererAddresses.value.addresses  | .[] | . == "'$hp'" ] | index(true)' config.json )
    #jq "del(.channel_group.values.OrdererAddresses.value.addresses[$index])"  config.json  > modified_config.json

    echo "==============application channel $CHANNEL_NAME remove orderer cert and orderer address======================="
    signNewOrderer

}


configPath="$GOPATH/src/bidchain/chaincode/fabric_config_test.json"
channelList=$(jq -r '.channels[] | .name' $configPath)
echo "channelList: $channelList"
#set +e
for ch in $channelList; do
  removeOrdererFromApplicationChannel $ch
  removeOrdererFromApplicationChannel "${ch}-test"
done
rm *.json *.pb
removeOrdererFromSystemChannel
echo "========= Successfully remove orderer: \"$newOrdererHost : $newOrdererPort\" =========== "

