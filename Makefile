BINARY_NAME=pkb
DIR=./...
VERSION ?= $(shell head -n 1 VERSION)

define ajv-docker
	docker run --rm -v "${PWD}":/repo weibeld/ajv-cli:5.0.0 ajv --spec draft2020
endef

define vhs-docker
	docker run --rm -v ${PWD}:/vhs --workdir /vhs ${BINARY_NAME}/vhs:local
endef

.PHONY: build
build:
	@CGO_ENABLED=0 go build -ldflags "-X github.com/tx3stn/pkb/cmd.Version=${VERSION}" -o ${BINARY_NAME} .

.PHONY: generate-gifs
generate-gifs: build
	@docker build --tag ${BINARY_NAME}-vhs:demo -f ./.docker/demo-gif.Dockerfile .
	@$(vhs-docker) /vhs/.scripts/gifs/demo.tape
	@$(vhs-docker) /vhs/.scripts/gifs/new-no-edit.tape
	@$(vhs-docker) /vhs/.scripts/gifs/edit.tape
	@$(vhs-docker) /vhs/.scripts/gifs/accessible-mode.tape

.PHONY: install
install: build
	@sudo cp ./${BINARY_NAME} /usr/local/bin/${BINARY_NAME}

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


.PHONY: test-e2e
test-e2e: build
	@docker build . -f .docker/bats-tests.Dockerfile -t pkb/bats:local
	@docker run --rm -it -v ${PWD}:/code pkb/bats:local bats --verbose-run --formatter pretty /code/.scripts/e2e-tests
