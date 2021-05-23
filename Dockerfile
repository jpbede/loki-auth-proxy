FROM alpine:3.13.5

LABEL org.opencontainers.image.source = "https://github.com/jpbede/loki-auth-proxy"

ENTRYPOINT ["/loki-auth-proxy", "l"]
COPY loki-auth-proxy /
