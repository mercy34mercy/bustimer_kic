FROM golang:1.20.2-buster as builder

WORKDIR /app

FROM golang:1.20.2-buster as builder

WORKDIR /app
ENV CGO_ENABLED=1
COPY go.mod go.sum ./
RUN go mod download

RUN apt-get update && apt-get install build-essential -y

COPY . .
RUN go build -a -ldflags '-linkmode external -extldflags "-static"'
FROM alpine
RUN mkdir /app
WORKDIR /app

COPY --from=builder /app/main ./main
COPY --from=builder /app/main ./gorm.db

RUN chmod 777 ./gorm.db

CMD ["./main"]