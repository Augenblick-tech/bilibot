FROM alpine:latest

WORKDIR /app

COPY BiliUpDynamicBot /app/BiliUpDynamicBot
COPY config.toml /app/config.toml

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone

ENTRYPOINT ["/app/BiliUpDynamicBot"]