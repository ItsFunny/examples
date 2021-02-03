#!/usr/bin/env bash


sudo ./clean.sh
docker-compose -f docker-compose-order.yaml -f docker-compose-all.yaml   --project-name multiorganization up -d

sleep 5
docker exec -it cli ./scripts/joinchannel.sh  2 5 2


#sleep 2
#docker cp cli:/opt/gopath/src/github.com/hyperledger/fabric/peer/demochannel.block ../
