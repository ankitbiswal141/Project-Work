!bin/bash

# Get and Install the packages
wget https://github.com/protobuf/releases//download/v23.0/protoc-23.0-linux_x86_64.zip
unzip protoc-23.0-linux-x86_64.zip -d protoc

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