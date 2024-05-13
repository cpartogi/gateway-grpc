run:
	go run main.go

init-app:
	go mod init gateway-grpc
	go mod tidy

proto-gen:
	protoc proto/*/*.proto --go_out=pb --go_opt=paths=source_relative --go-grpc_out=require_unimplemented_servers=false:pb --go-grpc_opt=paths=source_relative -I=proto --experimental_allow_proto3_optional	

