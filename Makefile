restart: down up

init: down-clear pull up-build

stop:
	docker-compose stop
up:
	docker-compose up -d
up-build:
	docker-compose up --build -d

down:
	docker-compose down

down-remove:
	docker-compose down --remove-orphans

down-clear:
	docker-compose down -v --remove-orphans

pull:
	docker-compose pull

build:
	docker-compose build
clean:
	docker system prune -a -f
	docker volume rm $(docker volume ls -q)