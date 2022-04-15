all: edpc.pb.go edpc_grpc.pb.go

edpc.pb.go edpc_grpc.pb.go: edpc.proto
	protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	edpc.proto
