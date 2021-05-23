FROM alpine:3.13.5
ENTRYPOINT ["/loki-auth-proxy", "l"]
COPY loki-auth-proxy /