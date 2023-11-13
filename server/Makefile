include .env

migrate:
ifndef C
	$(error C is undefined)
endif
	goose -dir ./migrations postgres "user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD) dbname=$(POSTGRES_DB) sslmode=disable" ${C} ${T}


migrate_new:
ifndef NAME
	$(error NAME is undefined)
endif
	goose -dir ./migrations create $(NAME) sql

docker:
	docker compose up

sqlc:
	docker run --rm -v ".:/src" -w /src sqlc/sqlc generate

mocks:
	mockgen -destination internal/users/mocks/mocks.go notik/internal/users Repo,Handler,Usecase
	mockgen -destination internal/pages/mocks/mocks.go notik/internal/pages Repo,Handler,Usecase

run:
	go run app/main.go
