FROM golang:1.19.3-buster

RUN mkdir /go/src/app

WORKDIR /go/src/app

ADD . /go/src/app

CMD [ "go","run","main.go" ]