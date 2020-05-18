#!/usr/bin/env bash


docker stop $(docker ps -qa)
docker rm $(docker ps -qa)
docker rmi $(docker images -q --filter reference=docker_redis_cluster*:*) --force