package proxy

import (
	"github.com/fasthttp/router"
	fasthttpprom "github.com/jpbede/fasthttp-prometheus-middleware"
	"github.com/rs/zerolog"
	"github.com/valyala/fasthttp"
	"go.bnck.me/loki-auth-proxy/pkg/authenticators"
)

// Proxy represents the loki proxy
type Proxy struct {
	Distributor   string
	QueryFrontend string
	Querier       string
	Authenticator authenticators.Authenticator
	Prometheus    bool

	logger              *zerolog.Logger
	distributorClient   *fasthttp.HostClient
	queryFrontendClient *fasthttp.HostClient
	querierClient       *fasthttp.HostClient
}

var basicAuthPrefix = []byte("Basic ")
var clientName = "loki-auth-proxy"

func New(distributor, queryFrontend, querier string, authenticator authenticators.Authenticator) *Proxy {
	return &Proxy{
		Distributor:   distributor,
		QueryFrontend: queryFrontend,
		Querier:       querier,
		Authenticator: authenticator,

		distributorClient: &fasthttp.HostClient{
			Addr: distributor,
			Name: clientName,
		},
		queryFrontendClient: &fasthttp.HostClient{
			Addr: queryFrontend,
			Name: clientName,
		},
	}
}

// ProxyRequest proxies the request to the backend servers
func (p *Proxy) ProxyRequest(proxyClient *fasthttp.HostClient) func(ctx *fasthttp.RequestCtx, username string) {
	return func(ctx *fasthttp.RequestCtx, username string) {
		if p.logger != nil {
			p.logger.Debug().
				Str("for-host", string(ctx.Request.Host())).
				Str("backend", proxyClient.Addr).
				Msg("Proxying request to backend")
		}

		// prepare fasthttp context
		p.prepareContext(ctx, proxyClient)

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
}

// Run starts listening on given address
func (p *Proxy) Run(listenAddress string, opts ...Option) error {
	// run options on proxy object
	for _, opt := range opts {
		opt(p)
	}

	// create router
	r := router.New()
	// distributor
	r.POST("/loki/api/v1/push", p.AuthHandler(p.ProxyRequest(p.distributorClient)))
	r.POST("/api/prom/push", p.AuthHandler(p.ProxyRequest(p.distributorClient)))
	// querier
	r.ANY("/loki/api/v1/tail", p.AuthHandler(p.WebsocketProxy()))
	r.ANY("/api/prom/tail ", p.AuthHandler(p.WebsocketProxy()))
	// query-frontend
	r.ANY("/loki/{path:*}", p.AuthHandler(p.ProxyRequest(p.queryFrontendClient)))
	r.ANY("/api/{path:*}", p.AuthHandler(p.ProxyRequest(p.queryFrontendClient)))
	handler := r.Handler

	// add prometheus endpoint when enabled
	if p.Prometheus {
		prom := fasthttpprom.NewPrometheus("loki_auth_proxy")
		prom.Use(r)
		handler = prom.Handler
	}

	// now listen and serve
	return fasthttp.ListenAndServe(listenAddress, handler)
}
