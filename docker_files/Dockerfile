FROM golang:1.23.4-alpine3.21

RUN apk update && apk add git

# sql-migrateをインストール
RUN go install github.com/rubenv/sql-migrate/...@latest
# sqlboilerをインストール
RUN go install github.com/volatiletech/sqlboiler/v4@latest
RUN go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest

RUN mkdir -p /go/src
WORKDIR /go/src
ADD ../src /go/src

RUN go mod tidy
