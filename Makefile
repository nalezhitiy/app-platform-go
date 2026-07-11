.SILENT:
.NOTPARALLEL:

## Settings
.DEFAULT_GOAL := help

## Colors
COLOR_RESET   = \033[0m
COLOR_INFO    = \033[32m
COLOR_COMMENT = \033[33m
COLOR_MAGENTA = \033[35m

export RUN_AS_USER=$(shell id -u)
export APP_COMMIT=$(shell git rev-parse --short HEAD)
export APP_BUILD_AT=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
include .env
export

guard-%:
	@ if [ "${${*}}" = "" ]; then \
		echo "Environment variable $* not set"; \
		exit 1; \
	fi

## Довідка
help:
	printf "${COLOR_COMMENT}Usage:${COLOR_RESET}\n"
	printf " make [target]\n\n"
	printf "${COLOR_COMMENT}Available targets:${COLOR_RESET}\n"
	awk '/^[a-zA-Z\-_0-9\.@]+:/ { \
	helpMessage = match(lastLine, /^## (.*)/); \
	if (helpMessage) { \
		helpCommand = substr($$1, 0, index($$1, ":")); \
		helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
		printf " ${COLOR_INFO}%-30s${COLOR_RESET} %s\n", helpCommand, helpMessage; \
	} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

## Запуск оточення
container@start:
	docker compose up -d
.PHONY: container@start

## Зупинка оточення
container@stop:
	docker compose down --remove-orphans
.PHONY: container@stop

## Рестарт оточення
container@restart: container@stop container@start
.PHONY: restart

## pull and build container
container@build:
	mkdir -p .go/pkg .go/cache
	RUN_AS_USER=${RUN_AS_USER} docker compose build --no-cache --pull
.PHONY: container@build

## status docker containers
status:
	docker compose ps
.PHONY: status

## Перегляд логів оточення
container@logs:
	docker compose logs -f
.PHONY: container@logs

## Shell (Backend)
container@shell-back: container@start
	docker compose exec backend sh
.PHONY: container@shell-back

container@run-back:
	docker compose run --rm backend sh
.PHONY: container@run-back