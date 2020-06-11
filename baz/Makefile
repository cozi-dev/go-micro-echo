
GOPATH:=$(shell go env GOPATH)

.PHONY: proto
proto:
    
	protoc --proto_path=.:${GOPATH}/src --go_out=. --micro_out=. proto/**/*.proto
    

.PHONY: build
build: proto

	go build -o baz-service *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t baz-service:latest
