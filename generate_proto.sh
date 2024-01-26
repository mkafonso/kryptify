#!/bin/bash

# generate protobuf files
protoc --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    proto/*.proto

# move generated files one level outside the current folder
mv pb/proto/* ./pb

# remove the pb folder
rm -rf pb/proto
