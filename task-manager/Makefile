.DEFAULT_GOAL := build
.PHONY:fmt vet build

fmt:
	go fmt ./...

vet:
	go vet ./...

build: vet
	go build -o tasks *.go

clean: tasks
	rm tasks
