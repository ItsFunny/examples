#!/bin/bash

export CHANNEL_NAME=$1
export ORDERER_ADDRESS=$2
export ORDERER_CA=$3
export NEW_ORG_NAME=$4
export NEW_ORG_MSPID=$5

. scripts/utils.sh.org3

#echo "========= Creating config transaction to add $NEW_ORG_NAME to network =========== "
# Fetch the config for the channel, writing it to config.json
fetchChannelConfig $ORDERER_ADDRESS $ORDERER_CA ${CHANNEL_NAME} config.json

# Modify the configuration to append the new org
jq -s '.[0] * {"channel_group":{"groups":{"Application":{"groups": {"'$NEW_ORG_MSPID'":.[1]}}}}}' config.json ./channel-artifacts/$NEW_ORG_NAME.json > modified_config.json


# Compute a config update, based on the differences between config.json and modified_config.json, write it as a transaction to org3_update_in_envelope.pb
createConfigUpdate ${CHANNEL_NAME} config.json modified_config.json $NEW_ORG_NAME"_update_in_envelope.pb" $NEW_ORG_MSPID

#echo "========= Config transaction to add $NEW_ORG_NAME to network created ===== "


#echo "Signing config transaction"
#signConfigtxAsPeerOrg 1 $NEW_ORG_NAME"_update_in_envelope.pb"
#peer channel signconfigtx -f $NEW_ORG_NAME"_update_in_envelope.pb"



#echo "========= Submitting transaction========= "
peer channel update -f $NEW_ORG_NAME"_update_in_envelope.pb" -c ${CHANNEL_NAME} -o $ORDERER_ADDRESS --tls --cafile ${ORDERER_CA} 2> error.log || (cat error.log; exit 0)


#echo "========= Config transaction to add $NEW_ORG_MSPID to network submitted! =========== "
echo "successfully add $NEW_ORG_MSPID to $CHANNEL_NAME"

