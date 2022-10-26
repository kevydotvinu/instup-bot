SHELL=/bin/bash
IMAGE=localhost/kevydotvinu/instup-bot
ENGINE?=podman
NAME=instup-bot

.PHONY: all
all: build run

.PHONY: build
build: main.go
	${ENGINE} build . --tag ${IMAGE}

.PHONY: run
run:
	${ENGINE} run --detach --name ${NAME} --rm --net host ${IMAGE}

.PHONY: logs
logs:
	${ENGINE} logs --follow ${NAME}

.PHONY: kill
kill:
	${ENGINE} kill ${NAME}
