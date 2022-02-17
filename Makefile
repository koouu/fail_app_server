up:
	docker-compose build && docker-compose up -d && docker-compose exec app go run main.go
down:
	docker-compose down
run:
	docker-compose up -d && docker-compose exec app go run main.go

