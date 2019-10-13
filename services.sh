#!/usr/bin/env bash

export GOOS=linux
export CGO_ENABLED=0

cd news;   go get;go build -o news-linux-amd64;   echo built `pwd`;cd ..
cd storage;go get;go build -o storage-linux-amd64;echo built `pwd`;cd ..

docker build -t services/newsservice newsservice/
docker service rm newsservice
docker service create --name=newsservice --replicas=1 --network=<NETWORK_NAME> -p=6767:6767 services/newsservice

docker build -t services/storageservice storageservice/
docker service rm storageservice
docker service create --name=storageservice --replicas=1 --network=<NETWORK_NAME> -p=6868:6868 services/storageservice
