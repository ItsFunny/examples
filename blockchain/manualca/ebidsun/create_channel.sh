#!/bin/bash

docker exec cli bash -c "echo \$CORE_PEER_LOCALMSPID > ./msp.txt"
msp=$(cat ./workspace/msp.txt)
echo "msp": $msp

configFile="$GOPATH/src/bidchain/chaincode/fabric_info_test.json"
channelList=$(jq -r '.MSP | .["'$msp'"] | .channels[] ' $configFile)
peerList=$(jq -r '.MSP | .["'$msp'"] | .peers[]' $configFile)
EBIDSUN_ALPHA="ebidsun-alpha"
echo "=====channelList:" $channelList "====="
echo "=====peerList:" $peerList "====="
for channelName in $channelList; do
	echo "=====generating channel[$channelName] config file====="
	if [[ $channelName == $EBIDSUN_ALPHA ]]; then
		configtxgen -profile BidchainChannel -outputCreateChannelTx ./channel-artifacts/channel.tx -channelID $EBIDSUN_ALPHA
	else
		profile=$(echo $channelName | sed 's/^\w\|\s\w/\U&/g')Channel
		configtxgen -profile $profile -outputCreateChannelTx ./channel-artifacts/$channelName.tx -channelID $channelName
	fi

	echo "=====updating anchor peers====="
	if [[ $channelName == $EBIDSUN_ALPHA ]]; then
		configtxgen -profile BidchainChannel -outputAnchorPeersUpdate ./channel-artifacts/${msp}anchors.tx -channelID $EBIDSUN_ALPHA -asOrg $msp
	else
		profile=$(echo $channelName | sed 's/^\w\|\s\w/\U&/g')Channel
		configtxgen -profile $profile -outputAnchorPeersUpdate ./channel-artifacts/${msp}_${profile}_Anchors.tx -channelID $channelName -asOrg $msp
	fi
done

docker exec cli scripts/create_channel_step.sh
