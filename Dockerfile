FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
COPY web ./web

EXPOSE 8080

CMD ["./main"] 