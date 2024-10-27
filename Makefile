.PHONY: run generate_migration migrate migrate_dev visualize_schema

run:
	cd cmd && air -c .air.toml

air-init:
	cd cmd && air init -c ../.air.toml

# ==============================================================================
# Go migrate postgresql

generate_migration:
	atlas migrate diff migration_name --dir "file://ent/migrate/migrations" --to "ent://ent/schema" --dev-url "docker://postgres/15/test?search_path=public"

migrate:
	atlas migrate apply --dir "file://ent/migrate/migrations" --url "postgres://postgres:admin@localhost:5432/test-ent?search_path=public&sslmode=disable"

visualize_schema:
	go run -mod=mod ariga.io/entviz ./path/to/ent/schema 

migrate_dev: generate_migration migrate

# ==============================================================================
# Tools commands

swaggo:
	echo "Starting swagger generating"
	swag init -g **/**/*.go