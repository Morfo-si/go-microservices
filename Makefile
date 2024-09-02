# Define variables
CONTAINER_NAME ?= petclinic-pg
DB_NAME ?= petclinic
POSTGRES_PASSWORD ?= postgres
POSTGRES_USERNAME ?= postgres
POSTGRES_PORT ?= 5432
IMAGE_NAME ?= postgres
DATA_FILE=dat/data.sql
SCHEMA_FILE=dat/schema.sql

# Default target
all: start-db apply-db-schema populate-db run

# Check for Docker or Podman
DOCKER_CMD=$(shell command -v docker 2> /dev/null || command -v podman 2> /dev/null)

# Ensure the container is running before applying the schema
apply-db-schema:
	@if [ -z "$(DOCKER_CMD)" ]; then \
	    echo "Error: Docker or Podman is not installed."; \
	    exit 1; \
	fi
	@sleep 5  # Wait a few seconds to ensure the container is fully up
	@$(DOCKER_CMD) exec -i $(CONTAINER_NAME) psql -U postgres -c "CREATE DATABASE $(DB_NAME);" || echo "Database $(DB_NAME) already exists."
	@$(DOCKER_CMD) cp $(SCHEMA_FILE) $(CONTAINER_NAME):/tmp/schema.sql
	@$(DOCKER_CMD) exec -i $(CONTAINER_NAME) psql -U $(POSTGRES_USERNAME) -d $(DB_NAME) -f /tmp/schema.sql
	@echo "Schema applied to $(DB_NAME)."
	@$(DOCKER_CMD) exec -i $(CONTAINER_NAME) rm /tmp/schema.sql

# Populate the PostgreSQL container with data
populate-db:
	@if [ -z "$(DOCKER_CMD)" ]; then \
	    echo "Error: Docker or Podman is not installed."; \
	    exit 1; \
	fi
	@$(DOCKER_CMD) cp $(DATA_FILE) $(CONTAINER_NAME):/tmp/data.sql
	@$(DOCKER_CMD) exec -i $(CONTAINER_NAME) psql -U $(POSTGRES_USERNAME) -d $(DB_NAME) -f /tmp/data.sql
	@echo "Schema applied to $(DB_NAME)."
	@$(DOCKER_CMD) exec -i $(CONTAINER_NAME) rm /tmp/data.sql

# Run the API service
run:
	go run main.go

# Start the PostgreSQL container
start-db:
	@if [ -z "$(DOCKER_CMD)" ]; then \
	    echo "Error: Docker or Podman is not installed."; \
	    exit 1; \
	fi
	@$(DOCKER_CMD) run -d \
	    --name $(CONTAINER_NAME) \
	    -e POSTGRES_PASSWORD=$(POSTGRES_USERNAME) \
	    -e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) \
	    -p $(POSTGRES_PORT):5432 \
	    $(IMAGE_NAME)

# Stop the PostgreSQL container if it's running
stop-db:
	@if [ -z "$(DOCKER_CMD)" ]; then \
	    echo "Error: Docker or Podman is not installed."; \
	    exit 1; \
	fi
	@$(DOCKER_CMD) stop $(CONTAINER_NAME) || echo "$(CONTAINER_NAME) is not running."

# Clean up the PostgreSQL container
clean-db: stop-db
	@if [ -z "$(DOCKER_CMD)" ]; then \
	    echo "Error: Docker or Podman is not installed."; \
	    exit 1; \
	fi
	@$(DOCKER_CMD) rm $(CONTAINER_NAME) || echo "$(CONTAINER_NAME) container is not available."

.PHONY: all apply-db-schema clean-db populate-db run start-db stop-db