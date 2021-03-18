#!/usr/bin/env bash

docker-compose --project-name demo -f ./network/order/docker-compose-order.yaml down

docker-compose --project-name demo -f ./network/peer/docker-compose-local-simple.yaml down

docker-compose -f ./network/tls/order-ca.yaml -f ./network/tls/order-ca.yaml -f ./network/tls/tls-ca.yaml down

rm -rf ./artifacts/*
rm -rf ./parent/crypto-config/*

rm -rf ./crypto-config/caOrganizations/*

rm -rf ./crypto-config/ordererOrganizations/*
#
rm -rf ./crypto-config/peerOrganizations/*

docker stop $(docker ps -qa --filter name=${DOMAIN}) && docker rm $(docker ps -qa --filter name=${DOMAIN})
docker rmi $(docker images -q --filter reference=dev-*:*)

