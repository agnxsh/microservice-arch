FRONT_END_BINARY=frontApp
BROKER_BINARY=brokerApp

## up: starts all the containers in the background without forcing build
up: 
	@echo "Starting Docker Images...."
	docker-compose up -d
	@echo "Docker images have started!"

## up_build: stops the docker-compose (if they're running), builds all projects and starts docker compose
up_build: build_broker
	@echo "Stopping docker images (if they're running)"
	docker-compose down
	@echo "Building (when required) and starting docker images...."
	docker-compose up --build -d
	@echo "Docker images are built and started!"

## down: stops the docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

## build_broker: builds the broker binary as a Linux executable
build_broker:
	@echo "Building broker binary..."
	cd ../broker-service && env GOOS=linux CGO_ENABLED=0 go build -o ${BROKER_BINARY} ./cmd/api
	@echo "Done!"


