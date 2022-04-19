#!/bin/sh

protoc -I=api/grpc/protobuf/ --go_out=api/grpc/protobuf --go-grpc_out=api/grpc/protobuf --dart_out=grpc:api/grpc/protobuf/dart/ api/grpc/protobuf/*.proto
protoc -I=api/grpc/protobuf/ --dart_out=grpc:api/grpc/protobuf/dart/ google/protobuf/timestamp.proto
