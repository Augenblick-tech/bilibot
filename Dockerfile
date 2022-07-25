FROM alpine:latest

WORKDIR /app

COPY BiliUpDynamicBot /app/BiliUpDynamicBot
COPY config.toml /app/config.toml

RUN apk add tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata

ENTRYPOINT ["/app/BiliUpDynamicBot"]