FROM golang:alpine3.18 AS builder
RUN apk add --no-cache --update gcc g++

WORKDIR /ig_server

COPY go.* .
COPY . .

RUN CGO_ENABLED=1 go build

FROM alpine:3.18

RUN apk add --no-cache tzdata
ENV TZ=Asia/Tokyo

WORKDIR /app

COPY --from=builder /ig_server/ig_server .
COPY static ./static

CMD ["/app/ig_server"]
