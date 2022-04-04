#!/bin/sh

protoc -I=api/grpc/cosmos-gov-grpc/ --go_out=api/grpc/cosmos-gov-grpc/go --go-grpc_out=api/grpc/cosmos-gov-grpc/go --dart_out=grpc:api/grpc/cosmos-gov-grpc/dart/ api/grpc/cosmos-gov-grpc/*.proto
