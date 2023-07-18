#!/bin/bash

BASE=$(dirname $0)

rm -f $BASE/*.pb.go

protoc \
  --go_out=./ \
  --go_opt=paths=source_relative \
  --go-grpc_out=./ \
  --go-grpc_opt=paths=source_relative \
  --go-grpc-mock_out=./ \
  --go-grpc-mock_opt=paths=source_relative \
  $BASE/example.proto
