FROM golang:1.26-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -ldflags "-X github.com/tx3stn/pkb/cmd.Version=e2e-test" -o pkb

FROM bats/bats:1.13.0

RUN apk add --no-cache \
	curl \
	musl-dev \
	expect

COPY --from=builder /app/pkb /usr/bin/pkb

ENTRYPOINT [ "bash" ]
