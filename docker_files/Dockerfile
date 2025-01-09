FROM golang:1.23.4-alpine3.21

RUN apk update && apk add git

RUN mkdir -p /go/src
WORKDIR /go/src
ADD ../src /go/src
