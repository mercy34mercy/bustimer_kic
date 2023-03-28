FROM golang:1.20.2-buster as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o main

FROM alpine
RUN mkdir /app
WORKDIR /app

COPY --from=builder /app/main ./main

CMD ["./main"]
