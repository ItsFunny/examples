#!/usr/bin/env bash

docker-compose -f ./network/tls/org-ca.yaml -f ./network/tls/order-ca.yaml -f ./network/tls/tls-ca.yaml up -d