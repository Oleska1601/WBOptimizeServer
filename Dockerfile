
FROM golang:1.25.8-alpine AS builder

WORKDIR /app

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY . .

RUN go build -o ./optimizer ./cmd/app/main.go

FROM alpine:latest AS runner

WORKDIR /app

COPY --from=builder /app/optimizer ./optimizer
COPY --from=builder /app/config/config.yaml ./config/config.yaml

EXPOSE 8081 8082 

CMD ["./optimizer"]
