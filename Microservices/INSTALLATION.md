## Here is all the commands required to run to accomodate the Accounts services:

-- Module 1:

1. wget https://github.com/protocolbuffers/protobuf/releases/download/v3.15.8/protoc-3.15.8-linux-x86_64.zip
unzip protoc-3.15.8-linux-x86_64.zip -d protoc
2. unzip protoc-23.0-linux-x86_64.zip -d protoc
3. sudo mv protoc/bin/protoc /usr/local/bin
4. go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
5. go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
6. echo $PATH
7. export PATH="$PATH:$(go env GOPATH)/bin"
8. source ~/.bashrc

>> Module 2 is available in /account section in ADD.md

>> Or you can RUN the INSTALL.sh directly, unless it causes some errors
-- cmd:
      chmod +x INSTALL.sh

      bash INSTALL.sh
      or
      ./INSTALL.sh
