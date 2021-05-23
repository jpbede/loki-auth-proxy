package proxy

import (
	"bytes"
	"encoding/base64"
	"github.com/valyala/fasthttp"
	proxy "github.com/yeqown/fasthttp-reverse-proxy/v2"
	"go.bnck.me/loki-auth-proxy/pkg/authenticators"
)

// Proxy represents the loki proxy
type Proxy struct {
	Backends      []string
	Authenticator authenticators.Authenticator
}

var basicAuthPrefix = []byte("Basic ")
var proxyServer *proxy.ReverseProxy

// proxyRequest proxies the request to the backend servers
func (p *Proxy) proxyRequest(ctx *fasthttp.RequestCtx, username string) {
	if proxyServer == nil {
		backendServers := map[string]proxy.Weight{}
		for _, backendServer := range p.Backends {
			backendServers[backendServer] = proxy.Weight(100 / len(p.Backends))
		}
		proxyServer = proxy.NewReverseProxy("", proxy.WithBalancer(backendServers))
	}

	// get org id form authenticator for username
	ctx.Request.Header.Add("X-Scope-OrgID", p.Authenticator.GetTenantID(username))
	proxyServer.ServeHTTP(ctx)
}

// AuthAndProxyHandler handler func for fasthttp that performs authentication and proxying
func (p *Proxy) AuthAndProxyHandler() func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		if auth := ctx.Request.Header.Peek("Authorization"); auth != nil {
			if bytes.HasPrefix(auth, basicAuthPrefix) {
				payload, err := base64.StdEncoding.DecodeString(string(auth[len(basicAuthPrefix):]))
				if err == nil {
					pair := bytes.SplitN(payload, []byte(":"), 2)
					if len(pair) == 2 && p.Authenticator.Authenticate(string(pair[0]), string(pair[1])) {
						p.proxyRequest(ctx, string(pair[0]))
						return
					}
				}
			}
		}

		// Request Basic Authentication otherwise
		ctx.Error(fasthttp.StatusMessage(fasthttp.StatusUnauthorized), fasthttp.StatusUnauthorized)
	}
}

// Run starts listening on given address
func (p *Proxy) Run(listenAddress string) error {
	return fasthttp.ListenAndServe(listenAddress, p.AuthAndProxyHandler())
}
