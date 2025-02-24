FROM golang:1.23.4-alpine AS builder

RUN apk update && apk add --no-cache git wget bash

RUN wget -qO- https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz \
    && mv migrate /usr/local/bin/

WORKDIR /app

ENV GOPROXY=direct
COPY ../migrations /app/migrations/

COPY ../go.mod ../go.sum ./
RUN go mod download

COPY ../configs/env/.env /app/.env
COPY ../. .

RUN go build -o main ./cmd/sm/main.go
COPY ../docs/ /docs/

FROM scratch AS runner

COPY --from=builder /app/main /main
COPY --from=builder /app/.env /configs/env/.env
COPY --from=builder /app/migrations /app/migrations
COPY --from=builder /docs /docs

WORKDIR /

ENTRYPOINT ["./main"]