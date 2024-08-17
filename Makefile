# Database connection string
DB_URL=postgres://postgres:root@localhost:5432/go-shorten?sslmode=disable

# Migrate create
migrate-create:
	@if [ -z "$(filter-out $@,$(MAKECMDGOALS))" ]; then \
		echo "Usage: make migrate-create <migration_name>"; \
		echo "Example: make migrate-create user_table"; \
		exit 1; \
	fi; \
	timestamp=$$(date +%Y%m%d); \
	migrate create -ext sql -dir migrations -seq $${timestamp}_$(filter-out $@,$(MAKECMDGOALS))

%:
	@:

# Migrate up
migrate-up:
	migrate -path migrations -database "$(DB_URL)" up

# Migrate down
migrate-down:
	migrate -path migrations -database "$(DB_URL)" down

# Migrate version
migrate-version:
	migrate -path migrations -database "$(DB_URL)" version

# Migrate force
migrate-force:
	@read -p "Enter version: " version; \
	migrate -path migrations -database "$(DB_URL)" force $$version

# Migrate to specific version
migrate-goto:
	@read -p "Enter version: " version; \
	migrate -path migrations -database "$(DB_URL)" goto $$version

# Migrate up by 1
migrate-up-1:
	migrate -path migrations -database "$(DB_URL)" up 1

# Migrate down by 1
migrate-down-1:
	migrate -path migrations -database "$(DB_URL)" down 1

.PHONY: migrate-create migrate-up migrate-down migrate-version migrate-force migrate-goto migrate-up-1 migrate-down-1