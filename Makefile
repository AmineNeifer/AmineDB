.PHONY: protobuf protoc-tools

protoc-tools:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.25.0
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0

protobuf:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		storepb/store.proto