# Build stage
FROM golang:1.20.2-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /app/main .

# Runtime stage
FROM alpine
RUN apk update && apk upgrade

RUN mkdir /app
WORKDIR /app

COPY --from=builder /app/main /app/main

CMD ["./main"]
