DATABASE_URL := postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable

.PHONY: test
test:
	go test -v ./...

.PHONY: certificate
certificate:
	cd ./cert && mkcert -key-file key.pem -cert-file cert.pem localhost

.PHONY: generate
generate:
	sqlc generate

.PHONY: migration
migration:
	migrate create -ext sql -dir db/migration -seq $(name)

.PHONY: migrate-up
migrate-up:
	migrate -path db/migrations -database $(DATABASE_URL) -verbose up

.PHONY: migrate-down
migrate-down:
	migrate -path db/migrations -database $(DATABASE_URL) -verbose down

.PHONY: migrate-drop
migrate-drop:
	migrate -path db/migrations -database $(DATABASE_URL) -verbose drop

.PHONY: migrate-force
migrate-force:
	migrate -path db/migrations -database $(DATABASE_URL) -verbose force "$(version)"

.PHONY: up
up:
	docker compose up --build

.PHONY: down
down:
	docker compose down