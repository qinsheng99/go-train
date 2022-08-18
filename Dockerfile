FROM golang:latest

MAINTAINER zjm

RUN mkdir /app/web

WORKDIR /app/web

COPY . /app/web
RUN cd /app/web

RUN go mod tidy && go build ./cmd/main.go

ENTRYPOINT ["/app/web/cmd/main"]