FROM golang:1.21 as builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=arm64 \
    GOBIN=/go/bin

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

# Install goose using a precompiled binary
RUN curl -L -o /go/bin/goose https://github.com/pressly/goose/releases/latest/download/goose_linux_x86_64 && \
    chmod +x /go/bin/goose

COPY . .

RUN go build -o main cmd/webserver/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/entrypoint.sh .
COPY --from=builder /app/db/migrations ./db/migrations
COPY --from=builder /go/bin/goose .

RUN chmod +x main entrypoint.sh /app/goose

EXPOSE 8080

CMD ["sh", "/app/entrypoint.sh"]
