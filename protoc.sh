#!/bin/sh

protoc -I=api/cosmos-gov-grpc/ --go_out=api/cosmos-gov-grpc/ --go-grpc_out=api/cosmos-gov-grpc/ --dart_out=grpc:api/cosmos-gov-grpc/dart/ api/cosmos-gov-grpc/*.proto
