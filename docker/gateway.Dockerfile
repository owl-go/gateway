FROM alpine:3.9.5

COPY ./bin/gateway /usr/bin/gateway
COPY ./configs/gateway.toml /usr/bin/config.toml

ENTRYPOINT ["/usr/bin/gateway"]