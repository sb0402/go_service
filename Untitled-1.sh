
go get -u google.golang.org/grpc
go get -u github.com/golang/protobuf/protoc-gen-go
vim ~/.bash_profile



cd /Users/0352mpind/Documents/projects/go_service/liquidparsing
protoc --go_out=. --go-grpc_out=. liquid.proto

