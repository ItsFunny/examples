#!/usr/bin/env bash


docker stop $(docker ps -qa)
docker rm $(docker ps -qa)
docker-compose -f docker-compose-org1.yaml  -f docker-compose-orderer.yaml   down --volumes --remove-orphans