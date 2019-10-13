#!/usr/bin/env bash

export GOOS=linux
export CGO_ENABLED=0

cd news;   go get;go run main.go;   echo built `pwd`;cd ..
cd storage;go get;go run storage.go;echo built `pwd`;cd ..