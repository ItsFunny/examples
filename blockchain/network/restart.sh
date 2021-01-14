#!/usr/bin/env bash


./clean.sh
docker-compose -f docker-compose.yaml  --project-name network up -d

docker exec -it cli ./scripts/joinchannel.sh