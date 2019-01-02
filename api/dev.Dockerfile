FROM golang

ENV GO111MODULE=on

RUN apt-get update && apt-get install inotify-tools -y

WORKDIR /app

COPY api/go.mod .
COPY api/go.sum .

COPY .env ../.env

RUN go mod download

COPY api/. .

#RUN go get github.com/cespare/reflex

EXPOSE 8080

COPY api/docker-entrypoint.sh /

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN go build -o ../api

ENTRYPOINT ["/docker-entrypoint.sh"]