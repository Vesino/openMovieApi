all:
	@echo "Available commands are:"
	@echo "- run-migrations-up   - Run migrations for postgres container, container should be up and running"
	@echo "- build_and_start  - Builds rest service image and starts docker cluster"
	@echo "- stop_remove_services  - Stops and remove containers services"

run-migrations-up:
	docker run --rm -v /Users/luisvargas/Desktop/Go/openMovieApi/migrations:/migrations --network host migrate/migrate -path=/migrations -database "postgres://greenlight:pa55word@localhost:7557/greenlight?sslmode=disable" up

build_and_start:
	docker-compose up --build

stop_remove_services:
	docker-compose down
