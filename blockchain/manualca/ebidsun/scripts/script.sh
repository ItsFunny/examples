#!/bin/bash

echo -e "Build your first network (BYFN) end-to-end test\n"

CHANNEL_NAME="$1"
DELAY="$2"
LANGUAGE="$3"
TIMEOUT="$4"
VERBOSE="$5"
NO_CHAINCODE="$6"
: ${CHANNEL_NAME:="ebidsun-alpha"}
: ${DELAY:="1"}
: ${LANGUAGE:="golang"}
: ${TIMEOUT:="10"}
: ${VERBOSE:="false"}
: ${NO_CHAINCODE:="false"}
LANGUAGE=`echo "$LANGUAGE" | tr [:upper:] [:lower:]`
COUNTER=1
MAX_RETRY=10

ORDERER_CA=/opt/workspace/crypto/ordererOrganizations/bidsun.com/orderers/orderer.bidsun.com/msp/tlscacerts/tlsca.bidsun.com-cert.pem

echo "Channel name : "$CHANNEL_NAME

# import utils
. scripts/utils.sh 2>/dev/null|| . utils.sh

createChannel() {
  setGlobals 0 1
  if [ -z "$CORE_PEER_TLS_ENABLED" -o "$CORE_PEER_TLS_ENABLED" = "false" ]; then
        set -x
	peer channel create -o orderer.bidsun.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/channel.tx >&log.txt
	res=$?
        set +x
  else
	set -x
	peer channel create -o orderer.bidsun.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/channel.tx --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA >&log.txt
	res=$?
	set +x
  fi
  cat log.txt
  verifyResult $res "Channel creation failed"
  echo -e  "===================== Channel '$CHANNEL_NAME' created =====================\n"
}

joinChannel () {
  for org in 1; do
    for peer in 0; do
	joinChannelWithRetry $peer $org
	echo "===================== peer${peer}.org${org} joined channel '$CHANNEL_NAME' ===================== "
	sleep $DELAY
	echo
    done
  done
}

## Create channel
echo "Creating channel..."
createChannel

## Join all the peers to the channel
echo "Having all peers join the channel..."
joinChannel

## Set the anchor peers for each org in the channel
echo "Updating anchor peers for org1..."
updateAnchorPeers 0 1

echo -e "\n========= All GOOD, BYFN execution completed =========== "
exit 0
