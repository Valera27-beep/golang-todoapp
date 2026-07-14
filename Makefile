include .env
export

export PROJECT_ROOT=$(shell pwd)

env-up:
	@docker compose up -d todoapp-postgres

env-down:
	@docker compose down

env-cleanup:
	read -p "Are you sure you want to remove the database volume? This action cannot be undone. (y/n): " ans && \
	if [ "$$ans" = "y" ]; then \
		docker compose down todoapp-postgres && \
		rm -rf out/pgdata; \
		echo "Files removed successfully."; \
	else \
		echo "Operation canceled."; \
	fi

migrate-create:
	@if [ -z "$$name" ]; then \
		echo "Error: provide migration name"; \
		exit 1; \
	fi
	migrate create -ext sql -dir migrations -seq $$name

migrate-up:
	make migrate-action action=up

migrate-down:
	make migrate-action action=down

migrate-action:
	@docker compose run --rm todoapp-postgres-migrate \
		-path /migrations \
		-database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@todoapp-postgres:5432/${POSTGRES_DB}?sslmode=disable" \
		$(action)

test-target:
	@echo "value: ${var}"

env-port-forward:
	@docker compose up -d port-forwarder

env-port-close:
	@docker compose down port-forwarder