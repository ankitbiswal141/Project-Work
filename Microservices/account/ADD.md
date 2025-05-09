1. option go_package = "github.com/<your github_repo_path>/{Project-Name}/account/pb;
2. RUN protoc --go_out=./pb --go-grpc_out=./pb account.proto