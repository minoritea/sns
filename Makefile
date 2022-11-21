.PHONY: proto-go proto-ts build-front init dev-server
init:
	$(MAKE) proto-go
	$(MAKE) build-front

./build:
	mkdir build

./bin:
	mkdir bin

./bin/buf: ./bin
	env GOBIN=$(CURDIR)/bin go install github.com/bufbuild/buf/cmd/buf@v1.9.0

./bin/protoc-gen-go: ./bin
	env GOBIN=$(CURDIR)/bin go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

./bin/protoc-gen-connect-go: ./bin
	env GOBIN=$(CURDIR)/bin go install github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go@latest

./build/proto: ./bin/buf ./bin/protoc-gen-go ./bin/protoc-gen-connect-go
	env PATH=$(CURDIR)/bin buf generate proto

./front/node_modules:
	cd front; npm install

./front/src/proto: ./front/node_modules
	cd front; npm run proto-ts

./build/front: ./front/node_modules ./front/src/proto
	cd front; npm run build

proto-go:
	rm -rf ./build/proto
	$(MAKE) ./build/proto

proto-ts: 
	rm -rf ./front/src/proto
	$(MAKE) ./front/src/proto

build-front:
	rm -rf ./build/front
	$(MAKE) ./build/front

dev-server:
	cd front; npm run dev -- --port 6501 &
	go run github.com/makiuchi-d/arelo@latest -p '**/*.go' -- go run .
