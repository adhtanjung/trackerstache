proto:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=require_unimplemented_servers=false:. --go-grpc_opt=paths=source_relative \
    pkg/pb/*.proto
server:
	go run cmd/main.go
