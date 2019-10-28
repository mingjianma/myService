#!/bin/bash
export GOOS=linux
export CGO_ENABLED=0

cd adsservice;go get;go build -o adsservice-linux-amd64;echo built `pwd`;cd ..

cd healthchecker;go get;go build -o healthchecker-linux-amd64;echo built `pwd`;cd ..

cp healthchecker/healthchecker-linux-amd64 adsservice/
docker build -t someprefix/adsservice adsservice/

docker service rm adsservice
docker service create --name=adsservice --replicas=1 --network=my_network -p=6767:6767 someprefix/adsservice