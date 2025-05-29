include .env
export $(shell sed 's/=.*//' .env)

.PHONY: migrations/up migrations/down db/start db/stop run

sqlc:
	sqlc generate

db/start:
	docker-compose up -d

db/stop:
	docker-compose down

migrations/up:
	docker run --rm -it \
		--network=host \
		-v "$(shell pwd)/db:/db" \
		ghcr.io/amacneil/dbmate --url=$(DATABASE_URL) up

migrations/down:
	docker run --rm -it \
		--network=host \
		-v "$(shell pwd)/db:/db" \
		ghcr.io/amacneil/dbmate --url=$(DATABASE_URL) down


migrations/new:
ifndef name
	$(error name is required: make migrations/new name=create_table)
endif
	docker run --rm -v "$(shell pwd)/db:/db" \
	  ghcr.io/amacneil/dbmate new $(name)

run:
	go run main.go

proto/gen:
	buf generate