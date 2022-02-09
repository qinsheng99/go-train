FROM golangweb

MAINTAINER z

RUN mkdir /app/web

WORKDIR /app/web

ADD . /app/web

RUN go build ./cmd/main.go

EXPOSE 111

CMD /app/web/cmd/main