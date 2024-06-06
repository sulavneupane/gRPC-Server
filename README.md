# gRPC Server
A simple project to setup a gRPC Server using Go language

## Install Proto Compiler (protoc) on MacOS
```shell
brew install protobuf-c
```

## Setup gPRC for GoLang
https://grpc.io/docs/languages/go/quickstart/

## Commands
### Compile proto files
```shell
protoc \
--go_out=invoicer \
--go_opt=paths=source_relative \
--go-grpc_out=invoicer \
--go-grpc_opt=paths=source_relative \
invoicer.proto
```
