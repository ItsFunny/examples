#!/usr/bin/env bash

docker stop $(docker ps -qa)
docker rm $(docker ps -qa)