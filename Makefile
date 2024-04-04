export DATABASE_URL := sqlite:.db/db.sqlite
export DBMATE_NO_DUMP_SCHEMA := true

dev:
	air

gen:
	sqlc generate

migrate:
	 mkdir -p .db && dbmate up

build:
	docker compose build

up:
	docker compose up -d

down:
	docker compose down

logs:
	docker compose logs -f
