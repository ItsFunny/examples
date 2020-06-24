#!/usr/bin/env bash


docker stop $(docker ps -qa)
docker rm $(docker ps -qa)
docker-compose -f docker-compose-cli.yaml -f docker-compose-orderer.yaml -f docker-compose-peer.yaml  down --volumes --remove-orphans