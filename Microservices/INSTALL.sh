!bin/bash

# Get and Install the packages
wget https://github.com/protocolbuffers/protobuf/releases/download/v3.15.8/protoc-3.15.8-linux-x86_64.zip
unzip protoc-3.15.8-linux-x86_64.zip -d protoc

# Move to the binary folder for command access
sudo mv protoc/bin/protoc /usr/local/bin

# Install Protobuff and GRPC for communication
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Check all the Paths
echo $PATH

# Set the Path
export PATH="$PATH:$(go env GOPATH)/bin"

#Refresh the kernel
source ~/.bashrc