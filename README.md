# Grafana Loki - Stateless multi tenant authentication proxy
[![PkgGoDev](https://pkg.go.dev/badge/go.bnck.me/loki-auth-proxy)](https://pkg.go.dev/go.bnck.me/loki-auth-proxy)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/a7da697dfae94a39bb74394ccf472df7)](https://www.codacy.com/gh/jpbede/loki-auth-proxy/dashboard)

Stateless multi tenant authentication proxy, that handles authentication and setting `X-Scope-OrgID` for Grafana Loki.

**work in progress**

## Requirements
To use this, Grafana Loki needs to be configured in multi tenant mode with `auth_enabled: true` as described in the [official docs](https://grafana.com/docs/loki/latest/operations/multi-tenancy/).

## Using it
An example configuration you can find at [example/loki-auth-proxy.yaml](example/loki-auth-proxy.yaml).
