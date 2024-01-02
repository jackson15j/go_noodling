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
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOBIN) v1.55.2

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
	docker build --progress=plain -t use-openapi-generated-server-testp --target=test .

run-docker: build-docker
	docker run -p 8083:8083 --rm --name petstore use-openapi-generated-server

ci-run-docker: build-docker
	docker run -d -p 8083:8083 --rm --name petstore use-openapi-generated-server
