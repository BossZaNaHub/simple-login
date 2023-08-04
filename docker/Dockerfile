FROM alpine

RUN addgroup -g 1000 -S kz \
    && adduser  kz -u 1000 -D -S -s /bin/bash -G kz \
    && mkdir -p /data \
    && chown -R kz:kz /data

RUN apk add --no-cache tzdata

USER 1000

COPY simple-login-linux-x64 /data/app/

WORKDIR /data/app

CMD ["./simple-login-linux-x64"]
