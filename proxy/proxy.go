package proxy

import (
	"bytes"
	"encoding/base64"
	"github.com/jpbede/loki-auth-proxy/authenticator"
	"github.com/jpbede/loki-auth-proxy/config"
	"github.com/valyala/fasthttp"
	proxy "github.com/yeqown/fasthttp-reverse-proxy/v2"
)

type Proxy struct {
	Config        *config.Config
	Authenticator authenticator.IAuthenticator
}

var basicAuthPrefix = []byte("Basic ")

func (p *Proxy) Run() error {
	backendServers := map[string]proxy.Weight{}
	for _, backendServer := range p.Config.Backends {
		backendServers[backendServer] = proxy.Weight(100 / len(p.Config.Backends))
	}
	proxyServer := proxy.NewReverseProxy("", proxy.WithBalancer(backendServers))

	return fasthttp.ListenAndServe(p.Config.HTTP.Listen, func(ctx *fasthttp.RequestCtx) {
		auth := ctx.Request.Header.Peek("Authorization")
		if bytes.HasPrefix(auth, basicAuthPrefix) {
			// Check credentials
			payload, err := base64.StdEncoding.DecodeString(string(auth[len(basicAuthPrefix):]))
			if err == nil {
				pair := bytes.SplitN(payload, []byte(":"), 2)
				if len(pair) == 2 && p.Authenticator.Authenticate(string(pair[0]), string(pair[1])) {
					ctx.Request.Header.Add("X-Scope-ID", p.Authenticator.GetTenantID(string(pair[0])))
					proxyServer.ServeHTTP(ctx)
				}
			}
		}

		// Request Basic Authentication otherwise
		ctx.Error(fasthttp.StatusMessage(fasthttp.StatusUnauthorized), fasthttp.StatusUnauthorized)
	})
}
