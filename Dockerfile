FROM alpine:3.13.5
COPY loki-auth-proxy /
EXPOSE 8081
VOLUME ["/etc/loki-auth-proxy.yaml"]
CMD ["/loki-auth-proxy", "-c", "/etc/loki-auth-proxy.yaml", "l"]
