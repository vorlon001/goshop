#!/usr/bin/bash

#go mod tidy

go build -o ./bin/clientgrpc ./cmd/client.go
go build -o ./bin/servergrpc ./cmd/server.go


docker build -t goshop-clientgrpc:v1 -f ./Dockerfile.clientgrpc .
docker build -t goshop-servergrpc:v1 -f ./Dockerfile.servergrpc .
