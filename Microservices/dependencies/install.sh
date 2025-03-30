#!bin/bash

# Update the system packages
sudo apt update

# Head to account/
cd ../account

# Install the additional packages
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Update the path
export PATH="$PATH:$(go env GOPATH)/bin"


