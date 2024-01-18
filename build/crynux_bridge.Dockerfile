FROM golang:alpine3.18 AS builder
RUN apk add --no-cache --update gcc g++

WORKDIR /crynux_bridge

COPY go.* .
COPY . .

RUN CGO_ENABLED=1 go build

FROM alpine:3.18

RUN apk add --no-cache tzdata
ENV TZ=Asia/Tokyo

WORKDIR /app

COPY --from=builder /crynux_bridge/crynux_bridge .
COPY static ./static

CMD ["/app/crynux_bridge"]
