# Variables
MIGRATE = migrate
DB_URL = mysql://root:root@tcp(localhost:3306)/multi_finance_golang_clean_architecture?charset=utf8mb4&parseTime=True&loc=Local
MIGRATION_DIR = database/migration
MAIN = cmd/web/main.go

# Targets
.PHONY: help create-migration run-migration run-app

help: ## Show this help message
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@awk 'BEGIN {FS = ":.*##"; printf "\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

create-migration: ## Create a new migration file
	@read -p "Enter migration name: " name; \
	$(MIGRATE) create -ext sql -dir $(MIGRATION_DIR) create_table_$$name


run-migration: ## Run all up migrations
	@read -p "Enter migration mode: " mode; \
	$(MIGRATE) -database "$(DB_URL)" -path $(MIGRATION_DIR) $$mode

run-app: ## Run the application
	go run $(MAIN)