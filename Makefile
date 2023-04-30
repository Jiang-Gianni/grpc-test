build:
	go build -o bin/app

run: build
	./bin/app

proto:
	protoc --go_out=. --go-grpc_out=.  --go_opt=paths=source_relative proto/service.proto

.PHONY: proto
