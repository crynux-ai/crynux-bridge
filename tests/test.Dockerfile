FROM golang:alpine3.18
RUN apk add --no-cache --update gcc g++

WORKDIR /app

COPY go.* .

RUN CGO_ENABLED=1 go mod download

COPY . .

RUN mkdir data
RUN mkdir data/images
RUN mkdir data/images/models
RUN mkdir data/logs

COPY ./tests/config.yml ./config/config.yml

CMD go test -timeout 2000s -p 1 ./...
