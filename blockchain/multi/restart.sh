#!/usr/bin/env bash
./clean.sh
docker-compose -f docker-compose-order.yaml -f docker-compose-all.yaml   --project-name multiorganization up -d
sleep 5
docker exec -it cli ./scripts/joinchannel.sh  2 5 2
docker cp cli:/opt/gopath/src/github.com/hyperledger/fabric/peer/demochannel1.block ../
            