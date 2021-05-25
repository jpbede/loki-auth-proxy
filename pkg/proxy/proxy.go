package proxy

import (
	"bytes"
	"encoding/base64"
	"github.com/Mnwa/fasthttprouter-prometheus"
	"github.com/fasthttp/router"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"go.bnck.me/loki-auth-proxy/pkg/authenticators"
)

// Proxy represents the loki proxy
type Proxy struct {
	Backend       string
	Authenticator authenticators.Authenticator
	Prometheus    bool
	logger        *zerolog.Logger
}

var basicAuthPrefix = []byte("Basic ")
var proxyClient *fasthttp.HostClient
var clientName = "loki-auth-proxy"

// proxyRequest proxies the request to the backend servers
func (p *Proxy) proxyRequest(ctx *fasthttp.RequestCtx, username string) {
	if proxyClient == nil {
		proxyClient = &fasthttp.HostClient{
			Addr: p.Backend,
			Name: clientName,
		}
	}

	if p.logger != nil {
		p.logger.Debug().
			Str("for-host", string(ctx.Request.Host())).
			Str("backend", proxyClient.Addr).
			Msg("Proxying request to backend")
	}

	// prepare fasthttp context
	p.prepareContext(ctx)

	// get org id form authenticator for username
	ctx.Request.Header.Add("X-Scope-OrgID", p.Authenticator.GetTenantID(username))

	// run request against backend
	if err := proxyClient.Do(&ctx.Request, &ctx.Response); err != nil {
		if p.logger != nil {
			p.logger.Error().
				Err(err).
				Msg("Errored while getting response from backend")
		}

		ctx.Response.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.Response.SetBody([]byte(err.Error()))
	} else if p.logger != nil {
		p.logger.Debug().
			Str("header", ctx.Response.Header.String()).
			Msg("Got response from backend")
	}

	// postprocess fasthttp context
	p.postProcessContext(ctx)
}

// AuthAndProxyHandler handler func for fasthttp that performs authentication and proxying
func (p *Proxy) AuthAndProxyHandler() func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		if p.logger != nil {
			p.logger.Debug().
				Str("host", string(ctx.Request.Host())).
				Msg("Got request to proxy, checking auth first")
		}

		if auth := ctx.Request.Header.Peek("Authorization"); auth != nil {
			if bytes.HasPrefix(auth, basicAuthPrefix) {
				payload, err := base64.StdEncoding.DecodeString(string(auth[len(basicAuthPrefix):]))
				if err == nil {
					pair := bytes.SplitN(payload, []byte(":"), 2)
					if len(pair) == 2 && p.Authenticator.Authenticate(string(pair[0]), string(pair[1])) {
						p.proxyRequest(ctx, string(pair[0]))
						return
					}
					if p.logger != nil {
						p.logger.Debug().
							Str("host", string(ctx.Request.Host())).
							Str("username", string(pair[0])).
							Msg("Auth invalid, rejecting")
					}
				} else if p.logger != nil {
					p.logger.Error().
						Str("host", string(ctx.Request.Host())).
						Err(err).
						Msg("A error occurred while checking auth")
				}
			}
		}

		hdl := fasthttpadaptor.NewFastHTTPHandler(promhttp.Handler())
		hdl(ctx)

		// Request Basic Authentication otherwise
		ctx.Error(fasthttp.StatusMessage(fasthttp.StatusUnauthorized), fasthttp.StatusUnauthorized)
	}
}

// Run starts listening on given address
func (p *Proxy) Run(listenAddress string, opts ...Option) error {
	// run options on proxy object
	for _, opt := range opts {
		opt(p)
	}

	// create router
	r := router.New()
	r.ANY("/{path:*}", p.AuthAndProxyHandler())
	handler := r.Handler

	// add prometheus endpoint when enabled
	if p.Prometheus {
		prom := fasthttprouter_prometheus.NewPrometheus("loki_auth_proxy")
		handler = prom.WrapHandler(r)
	}

	// now listen and serve
	return fasthttp.ListenAndServe(listenAddress, handler)
}
