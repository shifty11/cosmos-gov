#!/bin/sh

protoc -I=api/grpc/protobuf/ --go_out=api/grpc/protobuf/go --go-grpc_out=api/grpc/protobuf/go --dart_out=grpc:api/grpc/protobuf/dart/ api/grpc/protobuf/*.proto
