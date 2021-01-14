#!/usr/bin/env bash

docker stop $(docker ps -qa --filter name=democc)  && docker rm $(docker ps -qa --filter name=democc)
docker rmi $(docker images -qa --filter reference=dev-*:*)
docker-compose -f docker-compose.yaml  --project-name network down