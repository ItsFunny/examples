#!/bin/bash
set -eu

cleanChaincodeImage() {
  local CHAINCODE_IMAGE_NAME=$1
  echo "==================removing chaincodeName: $CHAINCODE_IMAGE_NAME start ===================="
  local INSTANTIATED_VERSION=$(docker exec cli peer chaincode list --instantiated -C ebidsun-alpha | grep "Name: $CHAINCODE_IMAGE_NAME," | awk -F, '{print $2}' | cut -d: -f2)
  INSTANTIATED_VERSION=${INSTANTIATED_VERSION:1}
  echo "instantiated version: $INSTANTIATED_VERSION"

  for chaincodeImageId in `docker images  | awk '{print $1}' | grep ^dev`
  do  
     local chaincodeImageName=$(echo $chaincodeImageId | awk -F- '{print $3}')
     local chaincodeImageVersion=$(echo $chaincodeImageId | awk -F- '{print $4}')
     if [ $chaincodeImageName = "" ]  || [ $chaincodeImageVersion = "" ]; then
        continue
     fi
     if [ "$chaincodeImageName" = "$CHAINCODE_IMAGE_NAME" ] && [ "$chaincodeImageVersion" != "$INSTANTIATED_VERSION" ];then
          echo "rm chaincode image id" $chaincodeImageId 
          docker rmi -f $chaincodeImageId 
      fi  
  done 
  echo "==================removing chaincodeName: $CHAINCODE_IMAGE_NAME end====================="
}

cleanAllChaincodeImages() {
   local chaincodeImageNames=$(docker images | grep ^dev | awk -F- '{print $3}' | sort | uniq)
   chaincodeImages=$(docker images  | awk '{print $1}' | grep ^dev)
   for image in $chaincodeImageNames
   do  
      cleanChaincodeImage $image
   done 
}

cleanAllChaincodeImages
