ARG GO_VERSION=1.23  # Valor padrão caso não seja passado

FROM golang:1.23 AS builder

WORKDIR /app

COPY --chown=root:root go.mod ./
RUN if [ -f go.sum ]; then cp go.sum ./; fi

RUN go mod tidy

COPY --chown=root:root . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/main ./cmd/server

FROM debian:stable-slim

WORKDIR /usr/local/bin

COPY --from=builder /app/bin/main /usr/local/bin/main

RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates tzdata && rm -rf /var/lib/apt/lists/*

ENV TZ=America/Sao_Paulo

EXPOSE 8080

LABEL env="production"

ENTRYPOINT ["/usr/local/bin/main"]
