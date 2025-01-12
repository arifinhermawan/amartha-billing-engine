.PHONY: build

# for dependencies
dep:
	@echo "RUNNING GO MOD TIDY..."
	@go mod tidy

	@echo "RUNNING GO MOD VENDOR..."
	@go mod vendor

docker-compose-up:
	@echo "Starting Docker containers using docker-compose"
	@docker-compose up -d

docker-compose-down:
	@echo "Stopping Docker containers using docker-compose"
	@docker-compose down

run: docker-compose-up
	@go run cmd/main.go

clean: docker-compose-down
	@docker stop amartha-billing-engine || true
