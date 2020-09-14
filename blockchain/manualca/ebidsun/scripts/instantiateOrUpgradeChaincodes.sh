#!/bin/bash
set -e

. scripts/utils.sh 2>/dev/null || . utils.sh

version_gt() {
	test "$(echo "$@" | tr " " "\n" | sort | head -n 1)" != "$1"
}

instantiateOrUpgradeChaincode() {
	PREFIX=$GOPATH/src/bidchain/chaincode
	CHAINCODE_PREFIX=bidchain/chaincode
	cd $PREFIX
	local index=0
	for chaincodePath in $chaincodePaths; do
		local file=$(basename $chaincodePath)
		CHANNEL_NAME=$(jq -r '.MSP | .["'$msp'"] | .chaincodes['$index'].channelName' $configPath)
		echo "CHANNEL_NAME:" $CHANNEL_NAME
		index=$((index + 1))

		CHAINCODE_PATH=$chaincodePath
		CHAINCODE_NAME=$file
		version=$(cat version_test.json | jq .$file.version)

		echo "chaincode path" $CHAINCODE_PATH
		version=${version:1:-1}

		if [ $version == "null" ] || [ $version == "" ]; then
			echo "invalid version $version for chaincode $CHAINCODE_NAME"
			exit 0
		fi

		# 低版本不更新
		INSTANTIATED_VERSION=$(peer chaincode list --instantiated -C $CHANNEL_NAME | grep "Name: $CHAINCODE_NAME," | awk -F, '{print $2}' | cut -d: -f2)
		echo "instantied version:" $INSTANTIATED_VERSION
		# 获取已经实例化链码版本
		instantiatedChaincodeVersion=$INSTANTIATED_VERSION
		# 低版本跳过
		if [[ "$INSTANTIATED_VERSION" != "" ]]; then
			if ! version_gt $version $INSTANTIATED_VERSION; then
				echo "skip instantitate or uprade chaincode $CHAINCODE_NAME, latest instantiated version: $INSTANTIATED_VERSION, wanted to instantiate or upgrade version: $version"
				echo "=================================================================================================================================="
				continue
			fi
		fi

		echo "chaincode version is" $version
		policy=$(cat version_test.json | jq .$file.policy)
		if [[ $policy == "null" ]] || [[ $policy == "" ]]; then
			policy=$(cat version_test.json | jq .default.policy)
		fi
		policy=${policy:1:-1}
		echo "policy is:" $policy

		# 之前没有实例化
		if [[ $instantiatedChaincodeVersion == "" ]]; then
			#实力化(初始化)操作
			echo "instantiated chaincode:" $file "version:" $version

			setGlobals 0 1
			startTime=$(date +%s)
			peer chaincode instantiate -o orderer.bidsun.com:7050 --tls --cafile $ORDERER_CA -C $CHANNEL_NAME -n $CHAINCODE_NAME -v $version -c '{"Args":[""]}' -P "$policy"
			deltaTime=$(($(date +%s) - startTime))
			ret=$?
			verifyResult $ret
			echo "succesfully instantiate chaincode" $CHAINCODE_NAME
			echo "instantiate chaincode time usage:" $deltaTime
		elif version_gt $version $instantiatedChaincodeVersion; then
			setGlobals 0 1
			startTime=$(date +%s)
			peer chaincode upgrade -o orderer.bidsun.com:7050 --tls --cafile $ORDERER_CA -C $CHANNEL_NAME -n $CHAINCODE_NAME -v $version -c '{"Args":[""]}' -P "$policy"
			endTime=$(date +%s)
			deltaTime=$(($endTime - startTime))
			ret=$?
			verifyResult $ret
			echo "succesfully upgrade chaincode" $CHAINCODE_NAME
			echo "upgrade chaincode time usage:" $deltaTime
			echo "endTime: " $endTime
		fi
		echo "=================================================================================================================================="
		sleep 1

	done
}

msp=$CORE_PEER_LOCALMSPID
echo "msp": $msp
configPath="$GOPATH/src/bidchain/chaincode/fabric_info_test.json"
chaincodePaths=$(jq -r '.MSP | .["'$msp'"] | .chaincodes[] | .chaincodePath' $configPath)
echo "chaincodes:" $chaincodePaths
#peerList=$(jq -r '.MSP | .["'$msp'"] | .peers[]' $configPath)
#echo "peers:" $peerList
chaincodeNum=$(echo $chaincodePaths | tr " " "\n" | wc -l)
echo "chaincodeNums: " $chaincodeNum
peerAddress="peer0.bidsun.com"
instantiateOrUpgradeChaincode $peerAddress
