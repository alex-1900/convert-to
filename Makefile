ifeq ($(wildcard .env.local),)
    include .env
else
    include .env.local
endif

RUN_COMMAND_PREFIX = docker run -it --rm -v "$$PWD":/usr/src/app -w /usr/src/app golang:1.22.4-alpine

.PHONY: dev build

mod_init:
	${RUN_COMMAND_PREFIX} go mod init github.com/alex-1900/convert-to
	${RUN_COMMAND_PREFIX} go mod tidy

sh:
	${RUN_COMMAND_PREFIX} sh

build:
	docker build -t ${PROJECT_NAME}:${VERSION} --target prod .

dev:
	docker build -t ${PROJECT_NAME}:dev --target runner .
	docker run -it ${PROJECT_NAME}:dev
