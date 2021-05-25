# Grafana Loki - Stateless multi tenant authentication proxy
[![PkgGoDev](https://pkg.go.dev/badge/go.bnck.me/loki-auth-proxy)](https://pkg.go.dev/go.bnck.me/loki-auth-proxy)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/a7da697dfae94a39bb74394ccf472df7)](https://www.codacy.com/gh/jpbede/loki-auth-proxy/dashboard)
[![codecov](https://codecov.io/gh/jpbede/loki-auth-proxy/branch/main/graph/badge.svg?token=g6JohOjiZm)](https://codecov.io/gh/jpbede/loki-auth-proxy)

Stateless multi tenant authentication proxy, that handles authentication and setting `X-Scope-OrgID` for Grafana Loki.

**Please be nice to me, the project is work in progress :)**

## Requirements
To use this, Grafana Loki needs to be configured in multi tenant mode with `auth_enabled: true` as described in the [official docs](https://grafana.com/docs/loki/latest/operations/multi-tenancy/).

## Installing

### Binaries
You will find pre-compiled binaries and packages for the most common OS under the [releases](https://github.com/jpbede/loki-auth-proxy/releases).

### macOS
Simply use `homebrew` (https://brew.sh/)

To install `loki-auth-proxy` use following command `brew install jpbede/tap/loki-auth-proxy`

## Using it
An example configuration you can find at [example/loki-auth-proxy.yaml](example/loki-auth-proxy.yaml).
