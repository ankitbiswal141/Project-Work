## This README is presented in the context the dependencies
## installation fails

# Here are the required commands to install the dependencies

-- Note: You need to change the directory to the account/

1. wget https://github.com/protocolbuffers/protobuf/releases/download/v23.0/protoc-23.0-linux-x86_64.zip
2. unzip protoc-23.0-linux-x86_64.zip -d protoc
3. sudo mv protoc/bin/protoc /usr/local/bin/
4. go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
5. go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
6. echo $PATH]
7. export PATH="$PATH:$(go env GOPATH)/bin"
8. source ~/.bashrc
9. create the pb folder in project
10. option go_package = "github.com/<PATH to your GitHub repository>"; /* This shall be added to your account.proto */
11. protoc --go_out=./pb --go-grpc_out=./pb.account.proto