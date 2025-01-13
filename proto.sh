#!/bin/bash
cd protobuf|| exit
protoc --proto_path=./ --go_out=../pkg/protobuf --go_opt=paths=source_relative --go-grpc_out=../pkg/protobuf --go-grpc_opt=paths=source_relative raidService.proto