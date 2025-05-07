BINARY_NAME=pkb
DIR=./...
VERSION ?= $(shell head -n 1 VERSION)

define ajv-docker
	docker run --rm -v "${PWD}":/repo weibeld/ajv-cli:5.0.0 ajv --spec draft2020
endef

.PHONY: build
build:
	@go build -ldflags "-X github.com/tx3stn/pkb/cmd.Version=${VERSION}" -o ${BINARY_NAME} .

.PHONY: install
install: build
	@sudo cp ./${BINARY_NAME} /usr/bin/${BINARY_NAME}

.PHONY: lint
lint:
	@golangci-lint run -v ${DIR}

.PHONY: schema-example-lint
schema-example-lint:
	@$(ajv-docker) validate -s /repo/schema/config.json -d /repo/schema/example.config.json

.PHONY: schema-validate
schema-validate:
	@$(ajv-docker) compile -s /repo/schema/config.json

.PHONY: test
test:
	@CGO_ENABLED=1 go test ${DIR} -race -cover

