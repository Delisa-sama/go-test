#!/usr/bin/env bash


docker service rm rabbitmq
docker build -t support/rabbitmq support/rabbitmq/
docker service create --name=rabbitmq --replicas=1 --network=<NETWORK_NAME> -p 1883:1883 -p 5672:5672 -p 15672:15672 support/rabbitmq