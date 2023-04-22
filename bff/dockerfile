FROM golang:1.20.2-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main .


FROM alpine
RUN apk update && apk upgrade
RUN mkdir /app
WORKDIR /app

ENV GO_ENV=prod
COPY --from=builder /app/main ./main
COPY ./gorm.db ./gorm.db

CMD ["./main"]