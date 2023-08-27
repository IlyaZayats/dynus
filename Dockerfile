FROM golang:1.19-alpine AS builder

WORKDIR /go/src/github.com/IlyaZayats/dynus
COPY . .

RUN go build -o ./bin/dynus ./cmd/dynus

FROM alpine:latest AS runner

COPY --from=builder /go/src/github.com/IlyaZayats/dynus/bin/dynus /app/dynus

RUN apk -U --no-cache add bash ca-certificates \
    && chmod +x /app/dynus

WORKDIR /app
ENTRYPOINT ["/app/dynus"]
