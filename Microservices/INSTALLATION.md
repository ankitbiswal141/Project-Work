1. wget https://github.com/protobuf/releases//download/v23.0/protoc-23.0-linux_x86_64.zip
2. unzip protoc-23.0-linux-x86_64.zip -d protoc
3. sudo mv protoc/bin/protoc /usr/local/bin
4. go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
5. go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
6. echo $PATH
7. export PATH="$PATH:$(go env GOPATH)/bin"
8. source ~/.bashrc

