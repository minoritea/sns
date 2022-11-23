SHELL := /usr/bin/env bash
.PHONY: proto-go proto-ts build-front init dev-server ssl-certs

init:
	$(MAKE) ssl-certs
	$(MAKE) proto-go
	$(MAKE) proto-ts
	$(MAKE) build-front

./build:
	mkdir build

./bin:
	mkdir bin

./bin/buf: | ./bin
	env GOBIN=$(CURDIR)/bin go install github.com/bufbuild/buf/cmd/buf@v1.9.0

./bin/protoc-gen-go: | ./bin
	env GOBIN=$(CURDIR)/bin go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

./bin/protoc-gen-connect-go: | ./bin
	env GOBIN=$(CURDIR)/bin go install github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go@latest

./bin/arelo: | ./bin
	env GOBIN=$(CURDIR)/bin go install github.com/makiuchi-d/arelo@latest

./bin/traefik: | ./bin
	curl -L https://github.com/traefik/traefik/releases/download/v2.9.5/traefik_v2.9.5_linux_amd64.tar.gz | tar zx -C ./bin traefik

./build/proto: | ./bin/buf ./bin/protoc-gen-go ./bin/protoc-gen-connect-go
	env PATH=$(CURDIR)/bin buf generate proto

./front/node_modules:
	cd front; npm install

./front/src/proto: ./front/node_modules
	cd front; npm run proto-ts

./ssl:
	mkdir -p ssl

./ssl/localhost.crt:
	mkdir -p ./ssl
	openssl req -x509 -out ssl/localhost.crt -keyout ssl/localhost.key \
  -newkey rsa:2048 -nodes -sha256 \
  -subj '/CN=localhost' -extensions EXT -config <( \
   printf "[dn]\nCN=localhost\n[req]\ndistinguished_name = dn\n[EXT]\nsubjectAltName=DNS:localhost\nkeyUsage=digitalSignature\nextendedKeyUsage=serverAuth")

./ssl/localhost.key:
	$(MAKE) ./ssl/localhost.crt

ssl-certs: ./ssl/localhost.crt ./ssl/localhost.key
	@:

proto-go:
	rm -rf ./build/proto
	$(MAKE) ./build/proto

proto-ts: 
	rm -rf ./front/src/proto
	$(MAKE) ./front/src/proto

dev-server: ./bin/traefik ./bin/arelo ./build/proto ./front/src/proto ./ssl/localhost.key ./ssl/localhost.crt ./front/node_modules
	cd front; npm run dev -- --port 6600 &
	bin/arelo -p '**/*.go' -- go run . &
	bin/traefik
