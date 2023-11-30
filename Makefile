export COMPOSE_PROJECT_NAME?=my-test-docker
export PORT?=5432

up:
	docker compose up -d
down:
	docker compose down
restart:
	docker compose down
	docker compose up -d

build:
	@go build -o .bin/app.exe cmd/store/main.go
run: build
	@.bin/app.exe