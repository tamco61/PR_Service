BINARY_NAME := myapp
DOCKER_IMAGE := $(shell basename $(CURDIR))
GO := go
DOCKER := docker
DOCKER_COMPOSE := docker-compose

.PHONY: all build run clean docker-build docker-up docker-down test mod-tidy

all: build

build:
	@echo "Сборка бинарника..."
	$(GO) build -o $(BINARY_NAME) ./main.go

run: build
	@echo "Запуск приложения..."
	./$(BINARY_NAME)

clean:
	@echo "Очистка..."
	rm -f $(BINARY_NAME)


mod-tidy:
	@echo "Обновление go.mod и go.sum..."
	$(GO) mod tidy

docker-build:
	@echo "Сборка Docker-образа: $(DOCKER_IMAGE)"
	$(DOCKER) build -t $(DOCKER_IMAGE) .

docker-up:
	@echo "Запуск через docker-compose..."
	$(DOCKER_COMPOSE) up --build -d

docker-down:
	@echo "Остановка docker-compose..."
	$(DOCKER_COMPOSE) down

restart: docker-down docker-up