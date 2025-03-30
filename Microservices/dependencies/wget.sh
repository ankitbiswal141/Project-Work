#!bin/bash

# Head to the required directory
cd ../account

# Updates the system packages
sudo apt update

# Command
wget https://github.com/protocolbuffers/protobuf/releases/download/v23.0/protoc-23.0-linux-x86_64.zip

# Make a clean up after the installation
clear

# Unzip the package
unzip protoc-23.0-linux-x86_64.zip -d protoc

# Move to the bin
sudo mv protoc/bin/protoc /usr/local/bin

# Make sure the packages are included in go mod
cd ..
go mod tidy


