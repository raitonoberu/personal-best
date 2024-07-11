export DATABASE_URL := sqlite:.db/db.sqlite

.PHONY: all

dev:
	air

gen:
	swag init && sqlc generate

migrate:
	 mkdir -p .db && dbmate up

rollback:
	dbmate rollback

s3:
	minio server .data

build:
	docker compose build

up:
	docker compose up -d

down:
	docker compose down

logs:
	docker compose logs -f

%:
	docker compose exec app ./main $@
