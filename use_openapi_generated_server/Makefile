# https://github.com/deepmap/oapi-codegen/blob/master/Makefile
GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin

help:
	@echo "This is a helper makefile for oapi-codegen"
	@echo "Targets:"
	@echo "    generate:    regenerate all generated files"
	@echo "    test:        run all tests"
	@echo "    gin_example  generate gin example server code"
	@echo "    tidy         tidy go mod"

$(GOBIN)/golangci-lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOBIN) v1.50.1

.PHONY: tools
tools: $(GOBIN)/golangci-lint

lint: tools
	$(GOBIN)/golangci-lint run ./...

generate:
	go generate ./...

test:
	go test -cover ./...

tidy:
	@echo "tidy..."
	go mod tidy

build-docker:
	docker build --progress=plain -t use-openapi-generated-server --target=app .

test-docker:
	docker compose up app_unittests --force-recreate -V --build
	docker compose --profile=ct up --force-recreate -V --build --exit-code-from=app_contracttests

run-docker: build-docker
	docker compose --profile=run up --force-recreate -V --build --exit-code-from=app

ci-run-docker: build-docker
	docker run -d -p 8087:8087 --rm --name petstore use-openapi-generated-server
