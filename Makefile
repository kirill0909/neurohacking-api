.PHONY:
.SILENT:


build:
	sudo docker-compose up -q --build  app

up:
	sudo docker-compose up -d

stop:
	docker-compose stop

ping:
	curl -k -X GET https://localhost:8000/

migrate:
	./migrate.sh
