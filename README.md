# Maackia

Maackia (マーキア) 犬槐 (山槐)

## Description

A gRPC Golang project.

## Preparation

### Configure protobuf

1. install protoc on your machine, ref to [releases](https://github.com/protocolbuffers/protobuf/releases)
2. install protoc-gen-go-grpc cmd
   1. go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
   2. go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
3. compile to `*.pb.go` files by using `protoc --go_out=. --go-grpc_out=. ./protos/*.proto`
