FROM golang:alpine as builder

RUN apk --no-cache add git

RUN apk --update add \
    go \
    musl-dev \
    util-linux-dev

WORKDIR /go/src/github.com/go/bilibot

COPY . .

RUN go build -a -ldflags "-linkmode external -extldflags '-static' -s -w" -o bilibot cmd/bilibot/main.go


FROM alpine:latest

WORKDIR /app

RUN apk add tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata

COPY --from=builder /go/src/github.com/go/bilibot/bilibot /app/bilibot

ENTRYPOINT ["/app/bilibot"]