#!/bin/bash

. scripts/utils.sh

configFile=$GOPATH/src/bidchain/chaincode/fabric_info_test.json
msp=$CORE_PEER_LOCALMSPID
channelList=$(jq -r '.MSP | .["'$msp'"] | .channels[] ' $configFile)
logFile=channel.log
peerList=$(jq -r '.MSP | .["'$msp'"] | .peers[]' $configFile)
EBIDSUN_ALPHA="ebidsun-alpha"
for channelName in $channelList; do
	echo "======creating channel:" $channelName "======"
	if [[ $channelName == "ebidsun-alpha" ]]; then
		set -x
		peer channel create -o orderer.bidsun.com:7050 -c $channelName -f ./channel-artifacts/channel.tx --tls true --cafile $ORDERER_CA &>$logFile
		ret=$?
		set +x
	else
		set -x
		peer channel create -o orderer.bidsun.com:7050 -c $channelName -f ./channel-artifacts/$channelName.tx --tls true --cafile $ORDERER_CA &>$logFile
		ret=$?
		set +x
	fi
	#echo "ret:" $ret
	if [[ $ret -ne 0 ]]; then
		grep -q "existing channel" $logFile
		if [[ $? -eq 0 ]]; then
			echo "channel $channelName already exists"
			continue
		else
			cat $logFile
			echo "程序异常终止!!!!"
			exit 0
		fi
	fi

	#set -ev

	echo "peerList" $peerList
	updateAnchorPeer=true
	for peerAddress in $peerList; do
		echo "===== $peerAddress Join channel[$channelName]====="
		CORE_PEER_ADDRESS=$peerAddress
		peer channel join -b $channelName".block" &>$logFile
		ret=$?
		if [[ $ret -ne 0 ]]; then
			updateAnchorPeer=false
			grep -q "LedgerID already exists" $logFile
			if [[ $? -ne 0 ]]; then
				cat $logFile
				exit 0
			fi
		fi
	done

	echo "=====updating anchor peers in channel[$channelName]====="
	#echo "updateAnchorPeer" $updateAnchorPeer
	if [[ $updateAnchorPeer == true ]]; then
		if [[ $channelName == "ebidsun-alpha" ]]; then
			set -x
			peer channel update -o orderer.bidsun.com:7050 -c $channelName -f ./channel-artifacts/${msp}anchors.tx --tls --cafile $ORDERER_CA
			ret=$?
			set +x
		else
			set -x
			profile=$(echo $channelName | sed 's/^\w\|\s\w/\U&/g')Channel
			peer channel update -o orderer.bidsun.com:7050 -c $channelName -f ./channel-artifacts/${msp}_${profile}_Anchors.tx --tls --cafile $ORDERER_CA
			ret=$?
			set +x
		fi
		#echo "ret:" $ret
		if [[ $ret -eq 0 ]]; then
			continue
		else
			exit 0
		fi
	fi
done
