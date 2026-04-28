TEMPLATE_FILE = compose-template.yaml
MAIN_COMPOSE = docker-compose.yaml

start: ## Start docker compose services
	@docker compose -f $(MAIN_COMPOSE) --env-file .env up -d

restart: ## Restart docker compose services
	@docker compose -f $(MAIN_COMPOSE) --env-file .env up -d --force-recreate

stop: ## Stop docker compose services
	@docker compose -f $(MAIN_COMPOSE) --env-file .env down

pull: ## Pull latest images needed by docker compose services
	@docker compose -f $(MAIN_COMPOSE) --env-file .env pull

clean: ## Stop docker composition and remove orphans
	@docker compose -f $(MAIN_COMPOSE) --env-file .env down --remove-orphans

generate: ## Generate a compose file for new service
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

help: ## Show this help message
	@which awk > /dev/null || (echo "awk not found. Please install it from https://www.gnu.org/software/gawk/manual/gawk.html" && exit 1)
	@echo ""
	@echo "This Makefile provides a set of commands to manage Docker Compose services."
	@echo "It allows you to start, stop, restart, pull, and generate new services."
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[$$()% a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-28s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
