FROM bats/bats:1.12.0

RUN apk add --no-cache \
	curl \
	musl-dev \
	expect

COPY pkb /usr/local/bin/pkb

ENTRYPOINT [ "bash" ]
