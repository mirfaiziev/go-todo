NO_COLOR=\033[0m
OK_COLOR=\033[32;01m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m

default: help

help: # Show this help
	@grep -E '^[a-zA-Z0-9 -]+:.*#'  Makefile | sort | \
		while read -r l; \
		 do printf "$(OK_COLOR)$$(echo $$l | cut -f 1 -d':')$(NO_COLOR)$$(echo $$l | cut -f 2- -d'#')\n"; \
	done


serve: # start the server
	go run cmd/todo/main.go

# create migration: migrate create -dir="./database/migrations" -ext="sql" table name
migrate: # Migtate up
	@echo "$(OK_COLOR)==> Running db migrations...$(NO_COLOR)"
	@migrate  -path="./database/migrations" -database "$(POSTGRES_DSN)" up

	