.PHONY:
.SILENT:

build:
	go build -o .bin/neurohacking-api cmd/main.go

run: build
	./run.sh

kill:
	./kill.sh

ping:
	curl -k -X GET https://localhost:8000/

dcbuild:
	docker-compose up -q --build db

up:
	docker-compose up db -d

stop:
	docker-compose stop

migrate:
	./migrate.sh
