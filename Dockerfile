FROM golang:1.20-bullseye AS base

    WORKDIR /usr/app

    # COPY go.mod go.sum ./
    COPY go.mod ./

    RUN go mod download -x

FROM base AS test

    COPY . ./

FROM test AS compiler

    RUN go build -o /bin/server ./cmd

FROM debian:stable-slim AS release

    RUN addgroup --gid 1000 --system appuser && \
        adduser --uid 1000 --gid 1000 --system appuser

    COPY --from=compiler --chown=appuser /bin/server /bin/server

    ENV SERVER_HTTP_PORT 8080

    EXPOSE ${SERVER_HTTP_PORT}

    ENV TINI_VERSION v0.18.0
    ADD https://github.com/krallin/tini/releases/download/${TINI_VERSION}/tini-static /bin/tini
    RUN chmod +x /bin/tini
    ENTRYPOINT ["/bin/tini", "--"]

    USER appuser

    CMD "/bin/server"