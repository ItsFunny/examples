#!/usr/bin/env bash

rm -rf /tmp/hyperledger/* && rm -rf /var/hyperledger/*
docker stop $(docker ps -qa --filter name=democc)  && docker rm -v $(docker ps -qa --filter name=democc)
docker rmi $(docker images -qa --filter reference=dev-*:*)
docker-compose -f docker-compose-order.yaml -f docker-compose-all.yaml  --project-name multiorganization down
if [[ -n ${1} ]];then
docker volume rm $(docker volume ls -qf dangling=true)
fi
