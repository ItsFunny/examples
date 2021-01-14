#!/bin/bash
# 1. 生成配置信息
cryptogen generate --config=crypto-config.yaml
# 2. 使用configtxgen 生成json文件
export FABRIC_CFG_PATH=${PWD} && configtxgen -printOrg Org6MSP > ./org6.json
# 3. 将这个peerOrganizations cp 到上级的crypto-config 里 并且将这个json文件挪到 artifacts中
cp -r ./crypto-config/peerOrganizations/org6.com ../../crypto-config/peerOrganizations/
cp ./org6.json ../../artifacts/

