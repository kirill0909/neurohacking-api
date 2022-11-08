.PHONY:
.SILENT:

build:
	go build -o .bin/neurohacking-api cmd/main.go

run:build
	./.bin/neurohacking-api
