#!/usr/bin/env bash

docker-compose --project-name demo -f ./network/order/docker-compose-order.yaml down

docker-compose --project-name demo -f ./network/peer/docker-compose-vlink-local-simple.yaml down

docker stop $(docker ps -qa --filter name=${DOMAIN}) && docker rm $(docker ps -qa --filter name=${DOMAIN})
docker rmi $(docker images -q --filter reference=dev-*:*)

docker-compose --project-name demo -f ./network/order/docker-compose-order.yaml up -d

docker-compose --project-name demo -f ./network/peer/docker-compose-vlink-local-simple.yaml up -d
