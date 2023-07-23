.PHONY: api mocks

CURRENT_USER=$(shell id -u):$(shell id -g)

api:
	go run ./cmd/api/main.go

mocks:
	docker run --rm -u $(CURRENT_USER) -v "$(PWD)":/src  -w /src vektra/mockery \
	--dir=internal/port/secondary \
	--name=TaskRepository \
	--filename=task_repository_mock.go \
	--output=internal/port/secondary/mocks \
	--with-expecter