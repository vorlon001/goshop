#!/usr/bin/bash

docker build -t docker.io/vorlon001/goshop:170820231311 -f ./Dockerfile .
docker build -t docker.io/vorlon001/goshop -f ./Dockerfile .

docker push docker.io/vorlon001/goshop:170820231311
docker push docker.io/vorlon001/goshop
