FROM golang:1.23.3-alpine AS builder

RUN apk update && apk add --no-cache git wget bash
WORKDIR /bsm
ENV GOPROXY=direct

COPY ../../go.mod ../../go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/main.go

FROM scratch AS runner
COPY --from=builder /bsm/main /main
COPY --from=builder /bsm/configs/env/.env /configs/env/.env
WORKDIR /

ENTRYPOINT ["./main"]