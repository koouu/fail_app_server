build:
	docker-compose build && docker-compose up -d && docker-compose exec app go run main.go
up:
	docker-compose up -d && docker-compose exec app go run main.go
down:
	docker-compose down
run:
	docker-compose exec app go run main.go
db:
	docker-compose exec mysql bash

