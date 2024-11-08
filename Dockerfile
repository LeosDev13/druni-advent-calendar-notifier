FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o app ./cmd/main.go

FROM alpine:latest

RUN apk add --no-cache libc6-compat

WORKDIR /app

COPY .env .env 

COPY --from=builder /app/app .

CMD ["./app"]
