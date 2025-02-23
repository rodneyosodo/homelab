TEMPLATE_FILE = compose-template.yaml
MAIN_COMPOSE = docker-compose.yaml

.PHONY: start
start:
	@docker compose -f $(MAIN_COMPOSE) --env-file .env up -d

.PHONY: restart
restart:
	@docker compose -f $(MAIN_COMPOSE) --env-file .env -d --force-recreate

.PHONY: stop
stop:
	@docker compose -f $(MAIN_COMPOSE) --env-file .env down

.PHONY: pull
pull:
	@docker compose -f $(MAIN_COMPOSE) --env-file .env pull

.PHONY: clean
clean:
	@docker compose -f $(MAIN_COMPOSE) --env-file .env down --remove-orphans

.PHONY: generate
generate:
	@if [ -z "$(service)" ]; then \
		echo "Usage: make generate service=<service_name>"; \
		exit 1; \
	fi
	@echo "Generating docker-compose file for service: $(service)"
	@mkdir -p $(service)
	@cp $(TEMPLATE_FILE) $(service)/docker-compose.yaml
	@sed -i "s/service_name/$(service)/g" $(service)/docker-compose.yaml
	@echo "  - path: ./$(service)/docker-compose.yaml" >> $(MAIN_COMPOSE)
	@echo "    project_directory: .." >> $(MAIN_COMPOSE)
	@echo "    env_file: .env" >> $(MAIN_COMPOSE)
	@echo "" >> $(MAIN_COMPOSE)
	@echo "Docker-compose file for $(service) generated"

.PHONY: help
help:
	@echo "This Makefile provides a set of commands to manage Docker Compose services."
	@echo "It allows you to start, stop, restart, pull, and generate new services."
	@echo ""
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  start                           - Start the services"
	@echo "  restart                         - Restart the services"
	@echo "  stop                            - Stop the services"
	@echo "  pull                            - Pull the latest images"
	@echo "  generate service=<service_name> - Generate a new service eg: make generate service=demo"
	@echo "  help                            - Display this help message"
