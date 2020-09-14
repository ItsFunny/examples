#!/bin/bash

# import utils
. scripts/utils.sh 2>/dev/null || . utils.sh

version_gt() {
	test "$(echo "$@" | tr " " "\n" | sort | head -n 1)" != "$1"
}

validateVersion() {
  local version=$1
  grep -Pq '^\d{1,3}(\.\d){1,2}$' <<<$version
  if [[ $? -ne 0 ]]; then
    echo "invalid version:" $version
    exit 0
  fi
}

installChaincodes() {
	PREFIX=$GOPATH/src/bidchain/chaincode
	cd $PREFIX
	for chaincodePath in $chaincodePaths; do
		for peerAddress in $peerList; do
			CORE_PEER_ADDRESS=$peerAddress
			local file=$(basename $chaincodePath)
			CHAINCODE_PATH=$chaincodePath
			CHAINCODE_NAME=$file
			version=$(cat version_test.json | jq .$file.version)

			echo "chaincode path" $CHAINCODE_PATH
			echo "chaincode name" $CHAINCODE_NAME
			version=${version:1:-1}

			if [ $version == "null" ] || [ $version == "" ]; then
				echo "invalid version $version for chaincode $CHAINCODE_NAME"
				continue
			fi
			validateVersion $version
			
			echo "chaincode version is" $version

			# 验证是否已经安装
			LATEST_INSTALLED_VERSION=$(peer chaincode list --installed | grep "Name: $CHAINCODE_NAME," | awk -F, '{print $2}' | awk -F: '{print $2}' | tr -d ' ' | tail -n 1)

			# 低版本跳过
			if [[ "$LATEST_INSTALLED_VERSION" != "" ]]; then
				if ! version_gt $version $LATEST_INSTALLED_VERSION; then
					echo "skip install chaincode $CHAINCODE_NAME, latest installed version: $LATEST_INSTALLED_VERSION, wanted to install version: $version"
					echo "========================================================================================"
					continue
				fi
			fi

			echo "Installing chaincode" $CHAINCODE_NAME "on $CORE_PEER_ADDRESS..."
            peer chaincode install -n $CHAINCODE_NAME -v $version -p $CHAINCODE_PATH 2> chaincode_install.log
            ret=$?
            if [[ $ret -ne 0 ]]; then
                cat chaincode_install.log
                exit 0
            fi 
			echo "succesfully install chaincode" $CHAINCODE_NAME "on $CORE_PEER_ADDRESS..."
            echo "========================================================================================"
		done
	done
}

configPath="$GOPATH/src/bidchain/chaincode/fabric_info_test.json"
msp=$CORE_PEER_LOCALMSPID
echo "========================================msp": $msp "======================================"
chaincodePaths=$(jq -r '.MSP | .["'$msp'"] | .chaincodes[] | .chaincodePath' $configPath | sort | uniq)
echo "chaincodeList:" $chaincodePaths
peerList=$(jq -r '.MSP | .["'$msp'"] | .peers[]' $configPath)
echo "peerList:" $peerList
installChaincodes
echo "=========================================="
