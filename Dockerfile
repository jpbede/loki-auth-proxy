FROM alpine:3.13.5
COPY loki-auth-proxy /
CMD ["/loki-auth-proxy", "-c", "/etc/loki-auth-proxy.yaml", "l"]
