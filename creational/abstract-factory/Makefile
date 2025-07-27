.PHONY: build up down clean

build:
	@echo -e "Building image...\n"
	docker build -t abstract-factory-go:latest .

up: build
	@echo -e "Launching services...\n"
	docker compose up -d

down:
	@echo -e "Shutting down the services...\n"
	docker compose down

clean:
	@echo -e "Stopping, removing all containers.\nRemoving networks.\nRemoving volumes.\n"
	docker compose down -v