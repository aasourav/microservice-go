BROKER_BINARY=brokerApp
AUTH_BINARY=authApp
LOGGER_BINARY=loggerApp

up:
	@echo "Starting docker images..."
	docker compose build --no-cache
	docker compose up -d
	# docker-compose up --build -d

up_build: build_broker
	@echo "Stopping docker images (if running...)"
	docker compose down
	@echo "Building (when required) and starting doker images"
	docker compose build --no-cache
	docker compose up -d
	@echo "Docker images built and started!"

up_build: build_auth
	@echo "Stopping docker images (if running...)"
	docker compose down
	@echo "Building (when required) and starting doker images"
	docker compose build --no-cache
	docker compose up -d
	@echo "Docker images built and started!"

up_build: build_logger
	@echo "Stopping docker images (if running...)"
	docker compose down
	@echo "Building (when required) and starting doker images"
	docker compose build --no-cache
	docker compose up -d
	@echo "Docker images built and started!"

down: 
	@echo "Stopping docker compose.."
	docker compose down
	@echo "Done.."

build_broker:
	@echo "building broker binary..."
	cd ../broker-service && env GOOS=linux CGO_ENABLED=0 go build -o ${BROKER_BINARY} ./cmd/api
	@echo "Done.."

build_auth:
	@echo "building auth binary..."
	cd ../authentication-service && env GOOS=linux CGO_ENABLED=0 go build -o ${AUTH_BINARY} ./cmd/api
	@echo "Done.."

build_logger:
	@echo "building log logger binary..."
	cd ../logger-service && env GOOS=linux CGO_ENABLED=0 go build -o ${LOGGER_BINARY} ./cmd/api
	@echo "Done.."