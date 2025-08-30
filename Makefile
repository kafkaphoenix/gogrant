.DEFAULT_GOAL := help

.PHONY: help
help: ## This help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: lint
lint: ## Run golangci-lint fixing issues
	golangci-lint run --fix

.PHONY: tests
tests: ## Run unit and integration tests
	go test ./... --tags=integration,unit -coverpkg=./...

.PHONY: swagger
swagger: ## Generate Swagger documentation available at http://localhost:8080/api/v0/swagger/index.html
	go tool swag init --parseDepth 1 -g ./cmd/app/main.go -o ./docs

.PHONY: mocks
mocks: ## Generate mocks
	go generate ./...

.PHONY: up
up: ## Start service
	docker-compose up -d

.PHONY: logs
logs: ## Show service logs
	docker-compose logs -f

.PHONY: clean
clean: ## clean docker containers, images, volumes and unused networks
	-docker rm -f `docker ps -a -q`
	-docker rmi -f `docker images -q`
	-docker volume rm `docker volume ls -q`
	docker network prune -f

# Testing
.PHONY: deploy
deploy: ## Deploy service from zero
	make clean
	go mod tidy
	make swagger
	make up
	make logs
