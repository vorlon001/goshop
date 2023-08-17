#!/usr/bin/bash

docker build -t docker.io/vorlon001/goshop-clientgrpc:170820231311 -f ./Dockerfile.clientgrpc .
docker build -t docker.io/vorlon001/goshop-servergrpc:170820231311 -f ./Dockerfile.servergrpc .

docker build -t docker.io/vorlon001/goshop-clientgrpc -f ./Dockerfile.clientgrpc .
docker build -t docker.io/vorlon001/goshop-servergrpc -f ./Dockerfile.servergrpc .

docker push docker.io/vorlon001/goshop-clientgrpc:170820231311
docker push docker.io/vorlon001/goshop-servergrpc:170820231311

docker push docker.io/vorlon001/goshop-clientgrpc
docker push docker.io/vorlon001/goshop-servergrpc
