BINARY_NAME=pkb
DIR=./...
VERSION ?= $(shell head -n 1 VERSION)

.PHONY: build
build:
	@go build -ldflags "-X github.com/tx3stn/pkb/cmd.Version=${VERSION}" -o ${BINARY_NAME} .

.PHONY: fmt
fmt:
	@go fmt ${DIR}

.PHONY: install
install: build
	@sudo cp ./${BINARY_NAME} /usr/bin/${BINARY_NAME}

.PHONY: lint
lint:
	@golangci-lint run -v ./...

.PHONY: lint-schema
lint-schema:
	@curl --request POST --header "Content-Type: application/json" --data @schema/config.json https://www.json-schema-linter.com/api/jsonschemalinter/analyze

.PHONY: push-tag
push-tag:
	@git tag -a ${VERSION} -m "Release ${VERSION}"
	@git push origin ${VERSION}

.PHONY: test
test:
	@CGO_ENABLED=1 go test ${DIR} -race -cover
